<?php

class proxyConnector {

    private static $instance;

    private $destinationUrl;
    private $userAgent;
    private $timeout;
    private $vector;
    private $url;
    private $returnData;
    private $ip;
    private $port;
    private $controlPort;
    private $controlPassword;
    private $curlError;
    private $switchIdentityAfterRequest = true;
    private $index = 0;

    public static function getIstance()
    {
        if (!isset(self::$instance)) {
            $c = __CLASS__;
            self::$instance = new $c;
        }

        return self::$instance;
    }

    public function setProxy($extIp="127.0.0.1", $extPort="9050")
    {
        $this->ip = $extIp;
        $this->port =$extPort;

    }

    public function setControlParameters($extPort, $extPassword)
    {
        $this->controlPassword = '"'.$extPassword.'"';
        $this->controlPort = $extPort;

    }

    public function getBinary($url, $dest)
    {
      $this->destinationUrl = str_replace(" ", "%20", $url);
      $this->curlError = false;
      $this->vector = null;
      $this->setUrl();

      $ch = curl_init($url);
      curl_setopt($ch, CURLOPT_PROXY, $this->ip .":". $this->port);
      curl_setopt($ch, CURLOPT_PROXYTYPE, CURLPROXY_SOCKS5);

      curl_setopt($ch, CURLOPT_URL, $this->url);
      curl_setopt($ch, CURLOPT_HEADER, 0);

      curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
      curl_setopt($ch, CURLOPT_BINARYTRANSFER, 1);

      curl_setopt($ch, CURLOPT_TIMEOUT, $this->timeout);

      $this->returnData = curl_exec($ch);
      $this->curlError = curl_error($ch);
      curl_close($ch);


      if(file_exists($dest)){
        unlink($dest);
      }

      $fp = fopen($dest, 'x');
      fwrite($fp, $this->returnData);
      fclose($fp);
    }

    public function launch($extUrl, $extVector, $fn = null, $postData = null, $extTimeout = null, $forbidden = 0)
    {
        //set parameters
        $this->destinationUrl = str_replace(" ","%20",$extUrl);
        $this->curlError = false;
        $this->vector = $extVector;

        if (empty($this->userAgent))
          $this->setUserAgent();

        //set url
        $this->setUrl();

        //if a timeout is set in the args, use it
        if(isset($timeout))
        {
            $this->timeout = $extTimeout;
        }

        //run cURL action against url
        $this->setCurl($fn, $postData);

        //renew identity
        if (($this->switchIdentityAfterRequest && ($this->index % 20) == 0) || $forbidden >= 2) {
            $this->newIdentity();
            $this->setUserAgent();
        }
	$this->index++;
    }

    public function getProxyData()
    {
        return array(
                'url' => $this->destinationUrl,
                'vector' => $this->vector,
                'userAgent' => $this->userAgent,
                'timeout' => $this->timeout,
                'proxy' => $this->ip .":". $this->port,
                'url' => $this->url,
                'curlError' => $this->curlError,
                'return' => $this->returnData
        );
    }

    public function newIdentity() {
            $fp = fsockopen($this->ip, $this->controlPort, $errno, $errstr, 30);
            if (!$fp) return false; //can't connect to the control port

            $auth_code = '"'.$this->controlPassword.'"'; // Old Password, not needed
            fputs($fp, "AUTHENTICATE $auth_code\r\n");

            $response = fread($fp, 1024);

            list($code, $text) = explode(' ', $response, 2);
            if ($code != '250') return false; //authentication failed

            //send the request to for new identity
            //            fputs($fp, "signal NEWNYM\r\n");
            // $response = fread($fp, 1024);
            // list($code, $text) = explode(' ', $response, 2);
            //if ($code != '250') return false; //signal failed

            // make it french
           fputs($fp, "SETCONF ExitNodes=2009f87590e626c98e1a5c8d08c23366c58b7951\r\n");
           echo "Changing Ip\n";
           $response = fread($fp, 1024);
           list($code, $text) = explode(' ', $response, 2);
           if ($code != '250') return false; //signal failed
           fclose($fp);
           return true;
    }

    public function loadDefaultSetUp() {

        $loaded_ini_array = parse_ini_file("./proxyConfiguration.ini",TRUE);

        $this->destinationUrl = null;
        $this->userAgent = null;
        $this->vector = null;
        $this->url = null;
        $this->returnData = null;

        $this->timeout = $loaded_ini_array["general"]["timeout"];
        $this->ip = $loaded_ini_array["general"]["ip"];
        $this->port = $loaded_ini_array["general"]["port"];
        $this->controlPort = $loaded_ini_array["TOR"]["controlPort"];
        $this->controlPassword = '"'.$loaded_ini_array["TOR"]["controlPassword"].'"';
        $this->switchIdentityAfterRequest = $loaded_ini_array["TOR"]["switchIdentityAfterRequest"];
    }


    ##PRIVATE
    private function  __construct() {
        $this->loadDefaultSetUp();
    }

    private function setUserAgent()
    {
        //list of browsers
        $agentBrowser = array(
                'Firefox',
                'Safari',
                'Opera',
                'Flock',
                'Internet Explorer',
                'Seamonkey',
                'Konqueror'
//                'GoogleBot'
        );
        //list of operating systems
        $agentOS = array(
                'Windows Seven',
                'Windows 8',
                'Windows 98',
                'Windows 2000',
                'Windows NT',
                'Windows XP',
                'Windows Vista',
                'Redhat Linux',
                'Ubuntu',
                'Fedora',
                'AmigaOS',
                'OS 10.5'
        );
        //randomly generate UserAgent
        $this->userAgent = $agentBrowser[rand(0,6)].'/'.rand(1,8).'.'.rand(0,9).' (' .$agentOS[rand(0,11)].' '.rand(1,7).'.'.rand(0,9).'; fr-FR;)';
    }


    private function setCurl($fn, $postData)
    {
      $ch = curl_init();

      $fpError = fopen('/tmp/curl_crawling-2.txt', 'w');
      $cookie_file = "/tmp/cookie.txt";

      curl_setopt($ch, CURLOPT_STDERR, $fpError);
      curl_setopt($ch, CURLOPT_PROXY, $this->ip);
      curl_setopt($ch, CURLOPT_PROXYPORT, $this->port);
      curl_setopt($ch, CURLOPT_PROXYTYPE, CURLPROXY_SOCKS5);
      curl_setopt($ch, CURLOPT_URL, $this->url);
      curl_setopt($ch, CURLOPT_HEADER, 0);

      if ($postData !== null) {
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLINFO_HEADER_OUT, true);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $postData);
        curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type: application/x-www-form-urlencoded; charset=UTF-8',
                                                   'Content-Length: 29',
                                                   'Origin: http://www.coursesu.com',
                                                   'Referer: http://www.coursesu.com/epicerie-salee/conserves/conserves/_pid3/136246',
                                                   'X-Prototype-Version: 1.7',
                                                   'X-Requested-With: XMLHttpRequest'
                                                   ));
      }

      curl_setopt($ch, CURLOPT_USERAGENT, $this->userAgent);
      curl_setopt($ch, CURLOPT_FOLLOWLOCATION, 1);
      curl_setopt($ch, CURLOPT_COOKIEFILE, $cookie_file);
      curl_setopt($ch, CURLOPT_COOKIEJAR, $cookie_file);
      curl_setopt($ch, CURLOPT_TIMEOUT, $this->timeout);

      if ($fn === null)
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
      else
        curl_setopt($ch, CURLOPT_FILE, $fn);

      $this->returnData = curl_exec($ch);
      if (preg_match("#forbidden#i", $this->returnData)) {
            $this->newIdentity();
            $this->setUserAgent();
	    echo "Renew, get 403..";
        }

      $this->curlError = curl_error($ch);
      $information = curl_getinfo($ch);
      fclose($fpError);
      curl_close($ch);
    }

    private function setUrl()
    {
        $this->url = $this->destinationUrl . $this->vector;
    }

    private function  __clone() {
        trigger_error("Clonig not allowed");
    }
}
