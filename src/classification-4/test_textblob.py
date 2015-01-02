from textblob.classifiers import NaiveBayesClassifier

with open('data/train_test.csv', 'r') as fp:
    cl = NaiveBayesClassifier(fp, format="csv")
    print cl.classify("Fleury Michon Hachis Parmentier, la barquette de 300 g")
