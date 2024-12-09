from flask import Flask, request
from dataHandler import handler
from DB import Database

app = Flask(__name__)

db = Database(debug=False)

@app.route('/', methods=['POST'])
def data():
    data = request.json
    handler(db, data, debug=False)
    return "200"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=27018)
