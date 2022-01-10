from flask import Flask, jsonify, request

app = Flask(__name__)

from 

@app.route('/products')
def getOrders():
    # return jsonify(products)
    return jsonify({'orders': products})