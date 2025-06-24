from flask import Flask, jsonify, request

app = Flask(__name__)

@app.route("/")
def index():
    return jsonify(status="ok", message="Service 2 Root")

@app.route("/ping")
def ping():
    return jsonify(status="ok", service="2")


@app.route("/hello")
def hello():
    print(f"➡️  /hello hit from {request.remote_addr}")
    return jsonify(message="Hello from Service 2")


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8002)
