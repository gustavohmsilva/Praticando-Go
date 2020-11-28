package dataparser

import (
	"os"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	testCase1, _ := os.Open("002.txt")
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
	}{
		{"Non-Existent TXT file", args{"002.txt"}, testCase1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile(t *testing.T) {
	testCase1, _ := Read("001.txt")
	testCase2, _ := Read("002.txt")
	p := Product{
		Name:     "Amazon Kindle PaperWhite (Nueva Version)",
		Category: "Dispositivos Amazon",
		Desc: `Te presentamos Kindle: ahora con luz frontal ajustable
para que puedas leer donde y cuando quieras. El Kindle está   
diseñado para la lectura y dispone de una pantalla táctil de
alto contraste que se lee como el papel impreso, sin ningún
reflejo, incluso bajo la luz del sol.`,
	}
	type args struct {
		f *os.File
	}
	tests := []struct {
		name    string
		args    args
		want    Product
		wantErr bool
	}{
		{"Test file 001.txt", args{testCase1}, p, false},
		{"Test file 002.txt", args{testCase2}, Product{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := File(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("File() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSV(t *testing.T) {
	testCase1, _ := Read("exported_products_black_friday.csv")
	testCase2, _ := Read("exported_prod...no_header.csv")
	testCase3and4, _ := Read("non_existent_file.csv")
	type args struct {
		f         *os.File
		hasHeader bool
	}
	tests := []struct {
		name    string
		args    args
		want    []Product
		wantErr bool
	}{
		{
			"Story Test Case",
			args{testCase1, true},
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
			},
			false,
		},

		{
			"No Header CSV file Test Case",
			args{testCase2, false},
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
			},
			false,
		},
		{
			"Story Test Case",
			args{testCase3and4, true},
			[]Product{},
			true,
		},
		{
			"Story Test Case",
			args{testCase3and4, false},
			[]Product{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CSV(tt.args.f, tt.args.hasHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("CSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
