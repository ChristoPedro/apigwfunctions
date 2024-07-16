from flask import Flask, Response

app = Flask(__name__)

@app.route("/rota1", methods=['GET'])
def rota1():
    return Response(response='{"Resposta": "Rota1"}',
                    mimetype="application/json") 

@app.route("/rota2", methods=['GET'])
def rota2():
    return Response(response='{"Resposta": "Rota2"}',
                    mimetype="application/json")

if __name__ == "__main__":
    
    app.run(debug=False, host='0.0.0.0')