# Work Shop OCI API Gateway + Funtions

Código parte do workshop de API Gateway + Functions.

Esse repositório contém 2 sample codes para demonstrar as capacidades do API gateway.

Além disso para as demonstrações também é utilizado como exemplo funções de auth do API Gateway que podem ser encontradas nesse repositório [aqui](https://github.com/ChristoPedro/authfunction).

- [Work Shop OCI API Gateway + Funtions](#work-shop-oci-api-gateway--funtions)
  - [OKE Functions](#oke-functions)
  - [Functions](#functions)
  - [Links Úteis](#links-úteis)
    - [API Gateway](#api-gateway)
    - [Functions](#functions-1)


## OKE Functions

O [oke-function](/oke-function/) consiste em um código python que utiliza o flask para export duas rotas, rota1 e rota2. Seu deploy é feito em um Cluster de Kubernetes através dos seguintes commandos:

```bash
kubectl run apigw --image pedrochristo/apigwpy

kubectl expose apigw --port 5000 --target-port 5000
```

Depois são criadas as rotas no ingress do Kubernetes:

```bash
kubectl create -f ingress.yaml
```

## Functions

O [functions](/functions/) consiste em um examplo de código para Functions Oracle escrito em GO. Cada um representa uma rota a ser consumida pelo API Gateway.

Para fazer o deploy basta ir em cada uma das paras e executar o comando:

```bash
fn deploy --app [Nome da Aplicação] -v
```

> Onde o **Nome da Aplicação** deve ser substituído pelo nome da sua aplicação criada no OCI.

## Links Úteis

### API Gateway

- [Quick Start](https://docs.oracle.com/en-us/iaas/Content/APIGateway/Tasks/apigatewayquickstartsetupcreatedeploy.htm)
  
- [Observabilidade](https://docs.oracle.com/en-us/iaas/Content/APIGateway/Tasks/apigatewayobservingapigatewaysandresources.htm)
  
- [Autenticação e Autorização](https://docs.oracle.com/en-us/iaas/Content/APIGateway/Tasks/apigatewayaddingauthzauthn.htm)
  
- [Parametrização](https://docs.oracle.com/en-us/iaas/Content/APIGateway/Tasks/apigatewaycallingandparameterizingapis.htm)
- [Cache](https://docs.oracle.com/en-us/iaas/Content/APIGateway/Tasks/apigatewayresponsecaching.htm)
  
- [Configurando Cache para usar o Serviço do OCI](https://medium.com/@bonnierockey/response-caching-for-oic-using-oci-api-gateway-and-oci-cache-with-redis-service-a987eab99f41)

### Functions

- [Fn Project](https://fnproject.io/)

- [Quick Start](https://docs.oracle.com/en-us/iaas/Content/Functions/Tasks/functionsquickstartguidestop.htm)

- [Pre-Build Functions](https://docs.oracle.com/en-us/iaas/Content/Functions/Tasks/functions_pbf_creating_prebuilt.htm)

- [Sample Functions](https://docs.oracle.com/en-us/iaas/Content/Functions/Tasks/functionsdownloadingsamples_topic-Sample_functions.htm#functionsdownloadingsamples_topic_Sample_functions)

- [Language Support](https://docs.oracle.com/en-us/iaas/Content/Functions/Tasks/languagessupportedbyfunctions.htm)

- [Trace](https://docs.oracle.com/en-us/iaas/Content/Functions/Tasks/functionstracing.htm)

- [Integrating](https://docs.oracle.com/en-us/iaas/Content/Functions/Tasks/functionsintegratingfunctions-toplevel.htm)