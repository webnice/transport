/*

package main

import (
	"io"
	"log"
	"time"

	"github.com/webdeskltd/transport"
)

func main() {
	// Отдельный объект для доступа к справочнику методов HTTP
	var tr = transport.NewTransport()

	// Создание запроса с полным набором необходимых атрибут
	req, err := tr.NewRequest(

		// Метод HTTP запроса ресурса из справочника
		tr.Method().Get()).

		// Адрес ресурса
		URL(`http://mtgjson.com/json/AllSets-x.json.zip`).

		// Referrer если требуется
		Referer(`http://mtgjson.com`).

		// UserAgent
		UserAgent(`Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.220 Safari/535.1`).

		// Content-Type
		ContentType(`application/zip`).

		// Максимальное время на запрос и полную загрузку ответа.
		// Если ответ будет загружаться дольше то соединение будет разорвано
		TimeOut(time.Minute * 30).

		// Для адресов https с самоподписанными сертификатами - отключение проверки SSL/TLS сеттификата сайта
		TLSVerifyOff().

		// Если результат необходимо записать во Writer, то передаём Writer
		// Если Writer не будет передан, то результат будет записан во временный файл,
		// который автоматически уничтожится в момент уничтожения объекта tr = transport.NewTransport() при сборки мусора runtime
		// Временные файлы созлаются во временной папке предоставляемой операционной системой
		//Response(writer).

		// Итоговый результат создания запроса возвращаем в error
		Error()

	if err != nil {
		log.Fatalf("Error create request: %s", err.Error())
		return
	}

	// Выполнение запроса
	// Writer не был передан, результат будет записан во временный файл
	rsp, err := req.Do()
	if err != nil {
		log.Fatalf("Error create request: %s", err.Error())
		return
	}

	// Не интересны результаты отличные от HTTP = 200
	if rsp.StatusCode() != 200 {
		log.Fatalf("Http error code is %d (%s)", rsp.StatusCode(), rsp.Status())
	}

	// Ожидаемый результат:
	// zip архив с одним файлом содержащим json текст который необходимо декодировать в переменную data
	// Для распаковки результат на лету устанавливаем Unzip(), который откроет контент как zip архив, считает первый файл и вернёт новый Content
	// Так как ожидается результат JSON, то контент декодируем Unmarshal в переменную data

	// Получение контента
	var data = make(map[string]Set)
	if err = rsp.Content().Unzip().UnmarshalJSON(&data); err != nil {
		log.Fatalf("Content error: %s", err.Error())
		return
	}

	log.Print("OK")

}

// Transform Трансформер для Content.Transform, пример функции
// Задача - Получать из Reader данные, преобразовывать их, предоставив на выходе Reader с преобразованными данными
// В случае не возможности предоставить Reader возвращается ошибка
// Применяется для фильтрации контента в цепочке. Пример:
// - Контент загружается, разархивируется zip, фильтруется через функцию Transform, декодируется XML в переменную responseData
// - Если контент не в кодировке UTF8 то при декодировании UnmarshalXML автоматически будет найдена кодировка контента и данные responseData будут в кодировке UTF-8
// - Для перекодирования используется golang.org/x/text/encoding, либо есть возможность указать свой перекодировщик с помощью rsp.Content().Transcode(encoding.Encoding)
// - Вызов:
// err = rsp.Content().Unzip().Transform(Transform).UnmarshalXML(&responseData)
// err = rsp.Content().Unzip().Transform(Transform).Transcode(myEncoder).UnmarshalXML(&responseData)
func Transform(rdr io.Reader) (ret io.Reader, err error) {
	ret = rdr
	return
}


*/

package transport
