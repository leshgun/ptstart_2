from flask import Flask, Response, jsonify

app = Flask(__name__)

@app.get("/")
def get_root() -> Response:
    hello_msg = "<h1>This is an example of a simple Rest Api server.</h1>"
    hello_msg += "<br><p>Available routes:</p>"
    hello_msg += "<ul style=\"list-style-type:square\">"
    hello_msg += "<li><a href=\"/healthz\">/healthz - check healthz.</a></li>"
    hello_msg += "</ul>"
    return hello_msg

@app.get("/healthz")
def get_healthz() -> Response:
    return jsonify({"code": 200, "status": "OK"})