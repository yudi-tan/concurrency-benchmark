from flask import Flask
import time
app = Flask(__name__)


def fibonacci(num):
    if num <= 1:
        return num
    return fibonacci(num-1) + fibonacci(num-2)


@app.route("/")
def hello():
    tmp = fibonacci(45)
    return str(tmp)


app.run(threaded=True)
