# IX - Lendo Arquivos Inteiramente ou Parcialmente

## O problema:   
Você finalmente foi aprovado para a posição para o qual fez a entrevista e sua primeira tarefa é começar a implementar um pequeno ***Projeto***, composto por várias partes, que a partir deste momento, chamaremos ***Stories***.   
Nesta primeira ***Story***, você terá de criar uma ou mais funções que leiam um arquivo de texto (.txt) formatado da seguinte forma:   
```
Nome do Produto   
Categoria do produto   
Descrição com múltiplas linhas descrevendo em texto plano   
o produto, suas aplicações e que tipos de indústrias mais   
se beneficiam com ele.   
```   
Cada arquivo sempre terá a mesma formatação, contendo na primeira linha uma informação, na segunda linha uma segunda informação e na terceira linha em diante uma série de informações em texto liso (sem negritos, itálicos, listas, enumerações e etc).


## Entrada de dados:    
Um arquivo de extensão .TXT     

## Saída de dados:   
Um tipo específico chamado Product

## Testes   
Entrada:   
Arquivo de exemplo neste folder   
Resultado esperado:  
```  
Product{
    Name: "Amazon Kindle PaperWhite (Nueva Version),
    Category: "Dispositivos Amazon",
    Desc: "Te presentamos Kindle: ahora con luz frontal ajustable   
    para que puedas leer donde y cuando quieras. El Kindle está   
    diseñado para la lectura y dispone de una pantalla táctil de   
    alto contraste que se lee como el papel impreso, sin ningún   
    reflejo, incluso bajo la luz del sol."
}  
```   