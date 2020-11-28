# X - Lendo Arquivos CSV (Comma-Separated Values)

## O problema:   
Nesta segunda ***Story*** do seu projeto para a empresa no qual foi contratado, o cliente escreveu que o package desenvolvido na ***Story*** anterior deve dar suporte não somente a arquivos de texto liso, como devem também dar suporte a arquivos em outros formatos. O formato de arquivo da vez é o [CSV](https://pt.wikipedia.org/wiki/Comma-separated_values). Abaixo é possível ver o ***mockup*** do formato CSV:   
   
```
NOME,CATEGORIA,DESCRICAO   
Produto 1,Categoria 1,"Descrição do Produto 1, contendo diversos tipos de caracteres, inclusive vírgulas!"
```   

## Entrada de dados:    
Um arquivo de extensão .CSV     

## Saída de dados:   
Uma relação de Product

## Testes   
### Entrada:   
Arquivo de exemplo neste folder   
### Resultado esperado:  
```  
[]Product{
		{
			Name:     "Logitech MX Master Ratón Inalámbrico",
			Category: "Teclados, ratones y periféricos",
			Desc:     "El ratón inalámbrico MX Master es el ratón estelar de Logitech, diseñado para ofrecer a los usuarios avanzados una comodidad, control y precisión superiores. Ofrece un atractivo diseño contorneado que se adapta a la mano y funciones desarolladas y está optimizado para Windows y Mac. MX Master cuenta con un botón rueda de precisión que cambia automáticamente del desplazamiento clic a clic al hiperrápido, un botón exclusivo para el pulsar para desplazamiento horizontal y otras funciones, el sensor láser Darkfield para un seguimiento perfecto (incluso sobre cristal) y una batería recargable con una duración de hasta 40 días. El ratón puede conectarse con el receptor Logitech Unifying incluido o mediante tecnología inalámbrica Bluetooth Smart. El molde del ratón se fabricó a partir de un diseño basado en una mano real. La duración de la batería depende del uso y del ordenador.",
		},
		{
			Name:     "Razer Klyo Webcam",
			Category: "Webcams y telefonía VoIP",
			Desc:     "La cámara Razer Kiyo es mucho más que una simple cámara de escritorio. Es un producto diseñado específicamente para crear el estudio de grabación y emisión perfecto. Tras analizar las necesidades de jugadores de todo el mundo, hemos creado un dispositivo que emite con una calidad increíble y una iluminación perfecta. Es potente, pero compacto. Además, es un complemento imprescindible para aquellos que desean dar el salto al mayor escenario del mundo.",
		},
		{
			Name:     "Wacom Bamboo Slate A5",
			Category: "Blocs y cuadernos de notas",
			Desc:     "Escribir notas y dibujar nunca ha sido tan fácil gracias al cuaderno inteligente Bamboo Folio, que permite trabajar de manera muy rápida combinando el placer de escribir en el papel tradicional con la practicidad y comodidad del digital. Gracias a la aplicación Wacom Inkspace, disponible para los dispositivos iOS, Android y Windows con Bluetooth, es posible convertir los apuntes en archivos digitales, sólo pulsando un bóton. El cuaderno digital puede guardar hasta 100 páginas aunque el dispositivo no esté conectado y sincronizarlas con servicios de nube como Inkspace, Dropbox, OneNote y Evernote.",
		},
	} 
```   