# I - Convertendo Milhas em Quilômetros

## O problema:   
Seu cliente lhe enviou uma documentação sobre as distâncias e consumos de combustível dos passeios turísticos que eles oferecem, porém existe um problema. A empresa é americana e eles enviaram todas as distâncias da documentação em Milhas. Por mais que você possa fazer esta conversão via calculadora do seu computador, ou até mesmo via Google, Nerd como você é, você decidiu escrever um pequeno programa que, ao receber um valor em milhas, o retorna em quilômetros.

## Entrada de dados:   
Valor em milhas, número

## Saída de dados:   
Valor em kilômetros

## Fórmula   
$$ Km = m * 1.609 $$

## Observação
Usei float64 em tudo, pois usando float32 tive problemas em arredondamento na hora dos testes.