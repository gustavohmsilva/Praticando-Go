# IV - Fábrica de Arruelas

## O problema:

Uma fábrica de material de construção gostaria de calcular o preço do frete por Kilograma das encomendas de arruelas fabricadas para cada cliente. Sendo este um cálculo bem complexo, diversas vezes os funcionários da empresa cometiam erros. Você foi contratado para solucionar este problema para a empresa, pois eles precisam de um programa simples para contabilizar com precisão este valor de frete.

## Entrada de dados:
Grossura da Arruela em centímetros, número   
Diâmetro externo da Arruela em centímetros, número   
Diâmetro interno da Arruela em centímetros, número   
Quantidade de Arruelas, número   
Valor do Frete por quilograma   

## Saída de dados:   

Valor total do frete (formatado para Reais)

## Valores imutáveis:   

Peso por volume do aço: 7,8 gramas por Centímetro Cúbico   


## Fórmulas

Volume do cilindro:   
$$
V = π * (D² - d²) / 4
$$


## Testes
Grossura da Arruela em centímetros: 0,2 cm    
Diâmetro externo da Arruela em centímetros: 2 cm     
Diâmetro interno da Arruela em centímetros: 1 cm   
Quantidade de Arruelas: mil (1000)   
Valor do Frete por quilograma: 2.50BRL   
Valor total da encomenda: "R$9.19"