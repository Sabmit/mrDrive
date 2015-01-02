import logging
import pandas as pd

from sklearn.feature_extraction.text import TfidfTransformer, CountVectorizer
from sklearn import metrics
from sklearn.pipeline import Pipeline
from sklearn.linear_model import SGDClassifier

###############################################################################
# Load some categories from the training set

df = pd.read_csv('data/train_test.csv', header=None, sep=',',
                         names=['products', 'categories'])

data_train = {};
data_train["data"] = df["products"][:14300]
data_train["target"] = df["categories"][:14300]

data_test = {};
data_test["data"] = df["products"][14300:]
data_test["target"] = df["categories"][14300:]

text_clf = Pipeline([('vect', CountVectorizer()),
                     ('tfidf', TfidfTransformer()),
                     ('clf', SGDClassifier(loss='hinge', penalty='l2', alpha=1e-5))])

text_clf = text_clf.fit(data_train["data"], data_train["target"])
predicted = text_clf.predict(data_test["data"])
print(metrics.classification_report(data_test["target"], predicted))

df_predict = pd.read_csv('data/data.csv', header=None, sep=',',
                        names=['products'])

predicted = text_clf.predict(df_predict["products"])
df = pd.DataFrame({'products' : df_predict["products"], 'categories' : predicted})

df.to_csv("result.csv", sep=',')
for doc, category in zip(df_predict["products"], predicted):
    print('%r => %s' % (doc, category))
