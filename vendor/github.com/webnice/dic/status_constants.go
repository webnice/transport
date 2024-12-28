// НЕ РЕДАКТИРОВАТЬ! Изменения будут перезаписаны при следующей кодогенерации.
// Code generated by go generate; DO NOT EDIT.

package dic

// Statuses Структура справочника статусов HTTP ответов.
type Statuses struct {
	// Unknown Статус: "", код: 0. Статус неизвестен.
	Unknown IStatus
	// Continue Статус: "Continue", код: 100. RFC9110, 15.2.1. Частично переданные на сервер данные проверены и сервер удовлетворён начальными данными. Клиент может продолжать отправку заголовков и данных.
	Continue IStatus
	// SwitchingProtocols Статус: "Switching Protocols", код: 101. RFC9110, 15.2.2. Сервер предлагает сменить протокол, перейти на более подходящий для указанного ресурса протокол. Список предлагаемых протоколов передаётся в заголовке Update Если клиент готов сменить протокол то новый запрос ожидается с указанием другого протокола.
	SwitchingProtocols IStatus
	// Processing Статус: "Processing", код: 102. RFC2518, 10.1. Запрос принят, на обработку запроса потребуется длительное время. Клиент, при получении такого ответа, должен сбросить таймер ожидания и ожидать следующего ответа в обычном режиме.
	Processing IStatus
	// EarlyHints Статус: "Early Hints", код: 103. RFC8297. Ранний возврат части заголовков, когда заголовки полного ответа не могут быть быстро сформированы.
	EarlyHints IStatus
	// NameNotResolved Статус: "Name Not Resolved", код: 105. Ошибка при разрешении доменного имени, в связи с не верным или отсутствующих ip адресом DNS сервера.
	NameNotResolved IStatus
	// OkToSendData Статус: "Ok To Send Data", код: 150. FTP. Согласие принимать данные от пользователя.
	OkToSendData IStatus
	// Ringing Статус: "Ringing", код: 180. Телефония. Уведомление о начале вызова на стороне вызываемого оборудования, соответствует длинному звуковому сигналу (КПВ) в телефонии.
	Ringing IStatus
	// Ok Статус: "OK", код: 200. RFC9110, 15.3.1. Успешный запрос. Результат запроса передаётся в заголовке или теле ответа.
	Ok IStatus
	// Created Статус: "Created", код: 201. RFC9110, 15.3.2. Успешный запрос, в результате которого был создан новый ресурс. В ответе возвращается заголовок Location с указанием созданного ресурса. Так же возвращаются характеристики нового ресурса, Content-Type.
	Created IStatus
	// Accepted Статус: "Accepted", код: 202. RFC9110, 15.3.3. Запрос принят на обработку, но обработка не завершена. Клиенту нет необходимости ожидать окончания обработки запроса, так как процесс может быть весьма длительным.
	Accepted IStatus
	// NonAuthoritativeInformation Статус: "Non-Authoritative Information", код: 203. RFC9110, 15.3.4. Успешный запрос, но передаваемая в ответе информация взята из кэша или не является гарантированно достоверной (могла устареть).
	NonAuthoritativeInformation IStatus
	// NoContent Статус: "No Content", код: 204. RFC9110, 15.3.5. Успешный запрос, в ответе передаются только заголовки без тела сообщения. Клиент не должен обновлять тело документа/ресурса, но может применить к нему полученные данные.
	NoContent IStatus
	// ResetContent Статус: "Reset Content", код: 205. RFC9110, 15.3.6. Сервер обязывает клиента сбросить введённые пользователем данные. Тела сообщения сервер клиенту не передаёт. Документ на стороне клиента обновлять не обязательно.
	ResetContent IStatus
	// PartialContent Статус: "Partial Content", код: 206. RFC9110, 15.3.7. Сервер удачно выполнил частичный GET запрос и возвращает только часть данных. В заголовке Content-Range сервер указывает байтовые диапазоны содержимого.
	PartialContent IStatus
	// MultiStatus Статус: "Multi-Status", код: 207. RFC4918, 11.1. Сервер возвращает результат выполнения сразу нескольких независимых операций. Результат ожидается в виде XML тела ответа с объектом multistatus.
	MultiStatus IStatus
	// AlreadyReported Статус: "Already Reported", код: 208. RFC5842, 7.1. Код возвращается в случае запроса к ресурсу состоящему из коллекций данных повторяющихся между собой и фактически является указателем на то что данные ресурса А можно взять из ресурса Б так как они идентичны, либо данные ответа текущего запроса с итерацией совпадают с ответом предыдущих запросов с итерацией (ссылка). Так же сервер возвращает заголовок вместе со ссылкой на данные на которые ссылается.
	AlreadyReported IStatus
	// ThisIsFine Статус: "This is fine", код: 218. Apache. Общая ошибка сервера, позволяющая передать тело ответа при включённой опции ProxyErrorOverride.
	ThisIsFine IStatus
	// ServerReady Статус: "Server Ready", код: 220. FTP. SMTP. Сервер готов обрабатывать запросы. Почтовый сервер готов к обслуживанию.
	ServerReady IStatus
	// ImUsed Статус: "IM Used", код: 226. RFC3229, 10.4.1. Заголовок A-IM от клиента был успешно принят и обработан. Сервер возвращает содержимое с учётом указанных параметров.
	ImUsed IStatus
	// LoginSuccessful Статус: "Login Successful", код: 230. FTP. Пользователь успешно подключился.
	LoginSuccessful IStatus
	// OkAcceptedAndProcessed Статус: "Ok Accepted And Processed", код: 250. SMTP. Команда принята и обработана.
	OkAcceptedAndProcessed IStatus
	// MultipleChoices Статус: "Multiple Choices", код: 300. RFC9110, 15.4.1. Существует несколько вариантов предоставления ресурса по типу MIME Сервер возвращает список альтернатив для выбора на стороне клиента.
	MultipleChoices IStatus
	// MovedPermanently Статус: "Moved Permanently", код: 301. RFC9110, 15.4.2. Запрошенный ресурс был окончательно перенесён на новый URI. Новый URI возвращается в заголовке Location.
	MovedPermanently IStatus
	// Found Статус: "Found", код: 302. RFC9110, 15.4.3. Запрошенный ресурс временно доступен по другому URI, возвращаемому в заголовке Location.
	Found IStatus
	// SeeOther Статус: "See Other", код: 303. RFC9110, 15.4.4. Ресурс по запрашиваемому URI необходимо запросить по адресу передаваемому в поле Location и только методом GET, несмотря на первоначальный запрос иным методом.
	SeeOther IStatus
	// NotModified Статус: "Not Modified", код: 304. RFC9110, 15.4.5. Сервер отвечает кодом 304, если ресурс был запрошен методом GET, с использованием заголовков If-Modified-Since или If-None-Match и документ не изменился с указанного момента. В этом ответе сервер не передаёт тело ресурса.
	NotModified IStatus
	// UseProxy Статус: "Use Proxy", код: 305. RFC9110, 15.4.6. Запрос к запрашиваемому ресурсу должен осуществляться через прокси URI, которого указывается в заголовки ответа в Location.
	UseProxy IStatus
	// Unused306 Статус: "Reserved Unused", код: 306. RFC9110, 15.4.7. Код зарезервирован и не используется.
	Unused306 IStatus
	// TemporaryRedirect Статус: "Temporary Redirect", код: 307. RFC9110, 15.4.8. Запрошенный ресурс, на короткое время, доступен по другому URI, указанному в заголовке ответа Location.
	TemporaryRedirect IStatus
	// PermanentRedirect Статус: "Permanent Redirect", код: 308. RFC9110, 15.4.9. Текущий запрос и все будущие запросы необходимо выполнить по другому URI ресурса. Новый адрес ресурса возвращается в заголовке Location, ответа сервера.
	PermanentRedirect IStatus
	// BadRequest Статус: "Bad Request", код: 400. RFC9110, 15.5.1. В запросе клиента присутствует синтаксическая ошибка или не указаны обязательные параметры запроса.
	BadRequest IStatus
	// Unauthorized Статус: "Unauthorized", код: 401. RFC9110, 15.5.2. Для доступа к ресурсу необходима аутентификация клиента. Заголовок ответа сервера будет содержать поле WWW-Authenticate, с перечнем условий аутентификации. Клиент может повторить запрос включив в новый запрос заголовок Authorization, с требуемыми для аутентификации данными.
	Unauthorized IStatus
	// PaymentRequired Статус: "Payment Required", код: 402. RFC9110, 15.5.3. Для доступа к ресурсу необходимо произвести оплату.
	PaymentRequired IStatus
	// Forbidden Статус: "Forbidden", код: 403. RFC9110, 15.5.4. Сервер отказывается выполнять запрос из-за ограничений в доступе клиента к указанному ресурсу.
	Forbidden IStatus
	// NotFound Статус: "Not Found", код: 404. RFC9110, 15.5.5. На сервере отсутствует запрашиваемый ресурс.
	NotFound IStatus
	// MethodNotAllowed Статус: "Method Not Allowed", код: 405. RFC9110, 15.5.6. Указанный клиентом HTTP метод запроса нельзя применять к запрашиваемому ресурсу. В ответе сервера будет заголовок Allow, с перечисленными через запятую методами запроса к ресурсу.
	MethodNotAllowed IStatus
	// NotAcceptable Статус: "Not Acceptable", код: 406. RFC9110, 15.5.7. Запрошенный ресурс не удовлетворяет переданным в заголовке характеристикам. Если запрос был не HEAD, тогда сервер в ответе вернёт список доступных характеристик, запрашиваемого ресурса.
	NotAcceptable IStatus
	// ProxyAuthenticationRequired Статус: "Proxy Authentication Required", код: 407. RFC9110, 15.5.8. Для доступа к ресурсу, необходима аутентификация клиента, на прокси сервере. Заголовок ответа сервера будет содержать поле WWW-Authenticate с перечнем условий аутентификации. Клиент может повторить запрос, включив в новый запрос заголовок Authorization с требуемыми для аутентификации данными.
	ProxyAuthenticationRequired IStatus
	// RequestTimeout Статус: "Request Timeout", код: 408. RFC9110, 15.5.9. Время ожидания сервера окончания передачи данных клиентом истекло, запрос прерван.
	RequestTimeout IStatus
	// Conflict Статус: "Conflict", код: 409. RFC9110, 15.5.10. Запрос не может быть выполнен из-за конфликта обращения к ресурсу, например, ресурс заблокирован другим клиентом.
	Conflict IStatus
	// Gone Статус: "Gone", код: 410. RFC9110, 15.5.11. Ресурс, ранее находившийся по указанному адресу, удалён и не доступен, серверу не известно новый адрес ресурса.
	Gone IStatus
	// LengthRequired Статус: "Length Required", код: 411. RFC9110, 15.5.12. Для выполнения запроса к указанному ресурсу, клиент должен передать Content-Length, в заголовке запроса.
	LengthRequired IStatus
	// PreconditionFailed Статус: "Precondition Failed", код: 412. RFC9110, 15.5.13. Сервер не смог распознать ни один заголовок или условие запроса, обычно используется совместно с If-Match.
	PreconditionFailed IStatus
	// RequestEntityTooLarge Статус: "Request Entity Too Large", код: 413. RFC9110, 15.5.14. Сервер отказывается выполнять запрос, из-за слишком большого тела запроса. В случае если проблема временная, тогда возвращается заголовок Retry-After с указанием времени, по истечении которого можно повторить запрос.
	RequestEntityTooLarge IStatus
	// RequestUriTooLong Статус: "Request URI Too Long", код: 414. RFC9110, 15.5.15. Сервер не может выполнить запрос к ресурсу из-за слишком большого запрашиваемого URI.
	RequestUriTooLong IStatus
	// UnsupportedMediaType Статус: "Unsupported Media Type", код: 415. RFC9110, 15.5.16. Сервер отказывается работать с указанным типом данных передаваемого контента, с использованием текущего метода HTTP запроса.
	UnsupportedMediaType IStatus
	// RequestedRangeNotSatisfiable Статус: "Requested Range Not Satisfiable", код: 416. RFC9110, 15.5.17. Не корректный запрос с указанием Range. В поле Range, заголовка запроса, указан диапазон за пределами размерности ресурса, и отсутствует поле If-Range. Если клиент передал байтовый диапазон, тогда сервер вернёт диапазоны ресурса, в заголовке Content-Range ответа.
	RequestedRangeNotSatisfiable IStatus
	// ExpectationFailed Статус: "Expectation Failed", код: 417. RFC9110, 15.5.18. Сервер не может удовлетворить переданный от клиента заголовок запроса Expect (ждать).
	ExpectationFailed IStatus
	// ImATeapot Статус: "I'm a teapot", код: 418. RFC9110, 15.5.19. RFC2324. Я чайник. Первоапрельская шутка от IETF :).
	ImATeapot IStatus
	// AuthenticationTimeout Статус: "Authentication Timeout", код: 419. Используется в качестве альтернативы коду 401, для запросов которые прошли проверку подлинности, но лишены доступа к определённым ресурсам сервера. Обычно код отдается, если CSRF-токен устарел или оказался некорректным.
	AuthenticationTimeout IStatus
	// EnhanceYourCalm Статус: "Enhance Your Calm", код: 420. Twitter. Ответ аналогичен ответу 429-слишком много запросов, призван заставить клиента отправлять меньшее количество запросов к ресурсу. В ответе сервера возвращается заголовок Retry-After, с указанием времени, по истечении которого можно повторить запрос к серверу.
	EnhanceYourCalm IStatus
	// MisdirectedRequest Статус: "Misdirected Request", код: 421. RFC9110, 15.5.20. Запрос был перенаправлен на сервер, не способный дать ответ.
	MisdirectedRequest IStatus
	// UnprocessableEntity Статус: "Unprocessable Entity", код: 422. RFC9110, 15.5.21. Не обрабатываемая сущность. Запрос к серверу корректный и верный, но в теле запроса имеется логическая ошибка, из-за которой не возможно выполнить запрос к ресурсу.
	UnprocessableEntity IStatus
	// Locked Статус: "Locked", код: 423. RFC4918, 11.3. В результате запроса ресурс успешно заблокирован.
	Locked IStatus
	// FailedDependency Статус: "Failed Dependency", код: 424. RFC4918, 11.4. Успешность выполнения запроса зависит от не разрешимой зависимости, другого запроса, который ещё не выполнен.
	FailedDependency IStatus
	// TooEarly Статус: "Too Early", код: 425. RFC8470, 5.2. Запрос коллекции ресурсов в не корректном порядке или запрос к упорядоченному ресурсу не по порядку.
	TooEarly IStatus
	// UpgradeRequired Статус: "Upgrade Required", код: 426. RFC9110, 15.5.22. Клиенту необходимо обновить протокол запроса. В заголовке ответа сервера Upgrade и Connection возвращаются указания, которые должен выполнить клиент для успешности запроса.
	UpgradeRequired IStatus
	// Unassigned Статус: "Not assigned", код: 427. WebDesk, Получены корректные данные запроса, но значения обязательных полей не установлены.
	Unassigned IStatus
	// PreconditionRequired Статус: "Precondition Required", код: 428. RFC6585, 3. Сервер требует от клиента выполнить запрос с указанием заголовков Range и If-Match.
	PreconditionRequired IStatus
	// TooManyRequests Статус: "Too Many Requests", код: 429. RFC6585, 4. Слишком много запросов от клиента к ресурсу. В ответе сервера возвращается заголовок Retry-After с указанием времени, по истечении которого, можно повторить запрос к серверу.
	TooManyRequests IStatus
	// RequestHeaderFieldsTooLarge Статус: "Request Header Fields Too Large", код: 431. RFC6585, 5. От клиента получено слишком много заголовков или длинна заголовка превысила допустимые размеры. Запрос прерван.
	RequestHeaderFieldsTooLarge IStatus
	// RequestedHostUnavailable Статус: "Requested Host Unavailable", код: 434. Запрашиваемый сервер ресурса не доступен.
	RequestedHostUnavailable IStatus
	// LoginTimeOut Статус: "Login Time-out", код: 440. Microsoft. Срок действия сеанса клиента истёк, и он должен снова войти в систему.
	LoginTimeOut IStatus
	// NoResponse Статус: "No Response", код: 444. Nginx. Сервер отказывается или не может вернуть результат запроса к ресурсу и незамедлительно закрыл соединение.
	NoResponse IStatus
	// RetryWith Статус: "Retry With", код: 449. Microsoft. В запросе к ресурсу не достаточно информации для его успешного выполнения. В заголовке ответа сервера передаётся Ms-Echo-Request с указанием необходимых полей.
	RetryWith IStatus
	// BlockedByParentalControls Статус: "Blocked by Parental Controls", код: 450. Microsoft. Запрос заблокирован системой 'родительский контроль'.
	BlockedByParentalControls IStatus
	// UnavailableForLegalReasons Статус: "Unavailable For Legal Reasons", код: 451. RFC7725, 3. Доступ к ресурсу закрыт по юридическим причинам или по требованиям органов государственной власти.
	UnavailableForLegalReasons IStatus
	// UnrecoverableError Статус: "Unrecoverable Error", код: 456. Обработка запроса вызывает не обрабатываемые сбои в базе данных или её таблицах.
	UnrecoverableError IStatus
	// RequestTerminated Статус: "Request Terminated", код: 487. Телефония. Инициатор вызова завершил попытку установить связь до соединения с вызываемым абонентом.
	RequestTerminated IStatus
	// RequestHeaderTooLarge Статус: "Request Header Too Large", код: 494. Nginx. Клиент отправил слишком большой запрос или слишком длинную строку заголовка.
	RequestHeaderTooLarge IStatus
	// SslCertificateError Статус: "SSL Certificate Error", код: 495. Nginx. Клиент предоставил недействительный SSL сертификат.
	SslCertificateError IStatus
	// SslCertificateRequired Статус: "SSL Certificate Required", код: 496. Nginx. Клиент не предоставил обязательный SSL сертификат.
	SslCertificateRequired IStatus
	// HttpRequestSentToHttpsPort Статус: "HTTP Request Sent to HTTPS Port", код: 497. Nginx. Клиент отправил HTTP-запрос на порт, прослушивающий HTTPS-запросы.
	HttpRequestSentToHttpsPort IStatus
	// InvalidToken Статус: "Invalid Token", код: 498. ArcGIS. Запрос выполнен с использованием токена с истекшим сроком действия или иным образом недействительным.
	InvalidToken IStatus
	// ClientClosedRequest Статус: "Client Closed Request", код: 499. Nginx. Клиент закрыл соединение не попытавшись получить от сервера ответ.
	ClientClosedRequest IStatus
	// InternalServerError Статус: "Internal Server Error", код: 500. RFC9110, 15.6.1. Любая не описанная ошибка на стороне сервера.
	InternalServerError IStatus
	// NotImplemented Статус: "Not Implemented", код: 501. RFC9110, 15.6.2. Сервер не имеет возможностей для удовлетворения доступа к ресурсу или реализация обработки запроса ещё не закончена.
	NotImplemented IStatus
	// BadGateway Статус: "Bad Gateway", код: 502. RFC9110, 15.6.3. Сервер, выступающий в роли шлюза или прокси, получил не корректный ответ от вышестоящего сервера.
	BadGateway IStatus
	// ServiceUnavailable Статус: "Service Unavailable", код: 503. RFC9110, 15.6.4. По техническим причинам, сервер не может выполнить запрос к ресурсу. В ответе сервер возвращает заголовок Retry-After, с указанием времени, по истечении которого, клиент может повторить запрос.
	ServiceUnavailable IStatus
	// GatewayTimeout Статус: "Gateway Timeout", код: 504. RFC9110, 15.6.5. Сервер, выступающий в роли шлюза или прокси, не дождался ответа от вышестоящего сервера.
	GatewayTimeout IStatus
	// HttpVersionNotSupported Статус: "HTTP Version Not Supported", код: 505. RFC9110, 15.6.6. Сервер отказывается поддерживать указанную в запросе версию HTTP протокола.
	HttpVersionNotSupported IStatus
	// VariantAlsoNegotiates Статус: "Variant Also Negotiates", код: 506. RFC2295, 8.1. Ошибка на стороне сервера связанная с циклической или рекурсивной задачей, которая не может завершиться.
	VariantAlsoNegotiates IStatus
	// InsufficientStorage Статус: "Insufficient Storage", код: 507. RFC4918, 11.5. Не достаточно места или ресурсов для выполнения запроса. Код может возвращаться как по причинам физической не хватки ресурсов, так и по причине текущих лимитов пользователя.
	InsufficientStorage IStatus
	// LoopDetected Статус: "Loop Detected", код: 508. RFC5842, 7.2. Бесконечный цикл. Запрос прерван, в результате вызванного на сервере бесконечного цикла, который не смог завершиться.
	LoopDetected IStatus
	// BandwidthLimitExceeded Статус: "Bandwidth Limit Exceeded", код: 509. Apache. Запрос прерван в связи с превышением, со стороны клиента, ограничений на скорость доступа к ресурсам.
	BandwidthLimitExceeded IStatus
	// NotExtended Статус: "Not Extended", код: 510. RFC2774, 7. На сервере отсутствует расширение, которое необходимо для успешного выполнения запроса к ресурсу.
	NotExtended IStatus
	// NetworkAuthenticationRequired Статус: "Network Authentication Required", код: 511. RFC6585, 6. Запрос был прерван сервером посредником, прокси или шлюзом, из-за необходимости авторизации клиента на сервере посреднике, до начала выполнения запросов.
	NetworkAuthenticationRequired IStatus
	// UnknownError Статус: "Unknown Error", код: 520. CloudFlare. Сервер CDN не смог обработать ошибку веб-сервера.
	UnknownError IStatus
	// WebServerIsDown Статус: "Web Server Is Down", код: 521. CloudFlare. Подключение сервера CDN отклоняются веб-сервером.
	WebServerIsDown IStatus
	// ConnectionTimedOut Статус: "Connection Timed Out", код: 522. CloudFlare. Серверу CDN не удалось подключиться к веб-серверу.
	ConnectionTimedOut IStatus
	// OriginIsUnreachable Статус: "Origin Is Unreachable", код: 523. CloudFlare. Веб-сервер недостижим для подключения с сервера CDN.
	OriginIsUnreachable IStatus
	// ATimeoutOccurred Статус: "A Timeout Occurred", код: 524. CloudFlare. Тайм-аут подключения между сервером CDN и веб-сервером истёк.
	ATimeoutOccurred IStatus
	// SslHandshakeFailed Статус: "SSL Handshake Failed", код: 525. CloudFlare. Ошибка рукопожатия SSL между сервером CDN и веб-сервером.
	SslHandshakeFailed IStatus
	// InvalidSslCertificate Статус: "Invalid SSL Certificate", код: 526. CloudFlare. Не удалось подтвердить сертификат шифрования веб-сервера.
	InvalidSslCertificate IStatus
	// RailgunError Статус: "Railgun Error", код: 527. CloudFlare. Соединение между CDN и веб-сервервером было прервано.
	RailgunError IStatus
	// NoSuchUserHere Статус: "No Such User Here", код: 550. SMTP. Ошибка: указанный почтовый ящик (пользователь) отсутствует.
	NoSuchUserHere IStatus
	// NetworkReadTimeoutError Статус: "Network Read Timeout Error", код: 598. Ответ возвращается сервером, который является прокси или шлюзом, перед вышестоящим сервером, и говорит о том, что сервер не смог получить ответ, на запрос, к вышестоящему серверу.
	NetworkReadTimeoutError IStatus
	// NetworkConnectTimeoutError Статус: "Network Connect Timeout Error", код: 599. Ответ возвращается сервером, который является прокси или шлюзом, перед вышестоящим сервером, и говорит о том, что сервер не смог установить связь или подключиться к вышестоящему серверу.
	NetworkConnectTimeoutError IStatus
	// Decline Статус: "Decline", код: 603. Телефония. Вызываемая сторона отклонила входящий вызов.
	Decline IStatus
}

func init() {
	singletonStatus = Statuses{
		Unknown:                       &tStatus{status: "", code: 0},
		Continue:                      &tStatus{status: "Continue", code: 100},
		SwitchingProtocols:            &tStatus{status: "Switching Protocols", code: 101},
		Processing:                    &tStatus{status: "Processing", code: 102},
		EarlyHints:                    &tStatus{status: "Early Hints", code: 103},
		NameNotResolved:               &tStatus{status: "Name Not Resolved", code: 105},
		OkToSendData:                  &tStatus{status: "Ok To Send Data", code: 150},
		Ringing:                       &tStatus{status: "Ringing", code: 180},
		Ok:                            &tStatus{status: "OK", code: 200},
		Created:                       &tStatus{status: "Created", code: 201},
		Accepted:                      &tStatus{status: "Accepted", code: 202},
		NonAuthoritativeInformation:   &tStatus{status: "Non-Authoritative Information", code: 203},
		NoContent:                     &tStatus{status: "No Content", code: 204},
		ResetContent:                  &tStatus{status: "Reset Content", code: 205},
		PartialContent:                &tStatus{status: "Partial Content", code: 206},
		MultiStatus:                   &tStatus{status: "Multi-Status", code: 207},
		AlreadyReported:               &tStatus{status: "Already Reported", code: 208},
		ThisIsFine:                    &tStatus{status: "This is fine", code: 218},
		ServerReady:                   &tStatus{status: "Server Ready", code: 220},
		ImUsed:                        &tStatus{status: "IM Used", code: 226},
		LoginSuccessful:               &tStatus{status: "Login Successful", code: 230},
		OkAcceptedAndProcessed:        &tStatus{status: "Ok Accepted And Processed", code: 250},
		MultipleChoices:               &tStatus{status: "Multiple Choices", code: 300},
		MovedPermanently:              &tStatus{status: "Moved Permanently", code: 301},
		Found:                         &tStatus{status: "Found", code: 302},
		SeeOther:                      &tStatus{status: "See Other", code: 303},
		NotModified:                   &tStatus{status: "Not Modified", code: 304},
		UseProxy:                      &tStatus{status: "Use Proxy", code: 305},
		Unused306:                     &tStatus{status: "Reserved Unused", code: 306},
		TemporaryRedirect:             &tStatus{status: "Temporary Redirect", code: 307},
		PermanentRedirect:             &tStatus{status: "Permanent Redirect", code: 308},
		BadRequest:                    &tStatus{status: "Bad Request", code: 400},
		Unauthorized:                  &tStatus{status: "Unauthorized", code: 401},
		PaymentRequired:               &tStatus{status: "Payment Required", code: 402},
		Forbidden:                     &tStatus{status: "Forbidden", code: 403},
		NotFound:                      &tStatus{status: "Not Found", code: 404},
		MethodNotAllowed:              &tStatus{status: "Method Not Allowed", code: 405},
		NotAcceptable:                 &tStatus{status: "Not Acceptable", code: 406},
		ProxyAuthenticationRequired:   &tStatus{status: "Proxy Authentication Required", code: 407},
		RequestTimeout:                &tStatus{status: "Request Timeout", code: 408},
		Conflict:                      &tStatus{status: "Conflict", code: 409},
		Gone:                          &tStatus{status: "Gone", code: 410},
		LengthRequired:                &tStatus{status: "Length Required", code: 411},
		PreconditionFailed:            &tStatus{status: "Precondition Failed", code: 412},
		RequestEntityTooLarge:         &tStatus{status: "Request Entity Too Large", code: 413},
		RequestUriTooLong:             &tStatus{status: "Request URI Too Long", code: 414},
		UnsupportedMediaType:          &tStatus{status: "Unsupported Media Type", code: 415},
		RequestedRangeNotSatisfiable:  &tStatus{status: "Requested Range Not Satisfiable", code: 416},
		ExpectationFailed:             &tStatus{status: "Expectation Failed", code: 417},
		ImATeapot:                     &tStatus{status: "I'm a teapot", code: 418},
		AuthenticationTimeout:         &tStatus{status: "Authentication Timeout", code: 419},
		EnhanceYourCalm:               &tStatus{status: "Enhance Your Calm", code: 420},
		MisdirectedRequest:            &tStatus{status: "Misdirected Request", code: 421},
		UnprocessableEntity:           &tStatus{status: "Unprocessable Entity", code: 422},
		Locked:                        &tStatus{status: "Locked", code: 423},
		FailedDependency:              &tStatus{status: "Failed Dependency", code: 424},
		TooEarly:                      &tStatus{status: "Too Early", code: 425},
		UpgradeRequired:               &tStatus{status: "Upgrade Required", code: 426},
		Unassigned:                    &tStatus{status: "Not assigned", code: 427},
		PreconditionRequired:          &tStatus{status: "Precondition Required", code: 428},
		TooManyRequests:               &tStatus{status: "Too Many Requests", code: 429},
		RequestHeaderFieldsTooLarge:   &tStatus{status: "Request Header Fields Too Large", code: 431},
		RequestedHostUnavailable:      &tStatus{status: "Requested Host Unavailable", code: 434},
		LoginTimeOut:                  &tStatus{status: "Login Time-out", code: 440},
		NoResponse:                    &tStatus{status: "No Response", code: 444},
		RetryWith:                     &tStatus{status: "Retry With", code: 449},
		BlockedByParentalControls:     &tStatus{status: "Blocked by Parental Controls", code: 450},
		UnavailableForLegalReasons:    &tStatus{status: "Unavailable For Legal Reasons", code: 451},
		UnrecoverableError:            &tStatus{status: "Unrecoverable Error", code: 456},
		RequestTerminated:             &tStatus{status: "Request Terminated", code: 487},
		RequestHeaderTooLarge:         &tStatus{status: "Request Header Too Large", code: 494},
		SslCertificateError:           &tStatus{status: "SSL Certificate Error", code: 495},
		SslCertificateRequired:        &tStatus{status: "SSL Certificate Required", code: 496},
		HttpRequestSentToHttpsPort:    &tStatus{status: "HTTP Request Sent to HTTPS Port", code: 497},
		InvalidToken:                  &tStatus{status: "Invalid Token", code: 498},
		ClientClosedRequest:           &tStatus{status: "Client Closed Request", code: 499},
		InternalServerError:           &tStatus{status: "Internal Server Error", code: 500},
		NotImplemented:                &tStatus{status: "Not Implemented", code: 501},
		BadGateway:                    &tStatus{status: "Bad Gateway", code: 502},
		ServiceUnavailable:            &tStatus{status: "Service Unavailable", code: 503},
		GatewayTimeout:                &tStatus{status: "Gateway Timeout", code: 504},
		HttpVersionNotSupported:       &tStatus{status: "HTTP Version Not Supported", code: 505},
		VariantAlsoNegotiates:         &tStatus{status: "Variant Also Negotiates", code: 506},
		InsufficientStorage:           &tStatus{status: "Insufficient Storage", code: 507},
		LoopDetected:                  &tStatus{status: "Loop Detected", code: 508},
		BandwidthLimitExceeded:        &tStatus{status: "Bandwidth Limit Exceeded", code: 509},
		NotExtended:                   &tStatus{status: "Not Extended", code: 510},
		NetworkAuthenticationRequired: &tStatus{status: "Network Authentication Required", code: 511},
		UnknownError:                  &tStatus{status: "Unknown Error", code: 520},
		WebServerIsDown:               &tStatus{status: "Web Server Is Down", code: 521},
		ConnectionTimedOut:            &tStatus{status: "Connection Timed Out", code: 522},
		OriginIsUnreachable:           &tStatus{status: "Origin Is Unreachable", code: 523},
		ATimeoutOccurred:              &tStatus{status: "A Timeout Occurred", code: 524},
		SslHandshakeFailed:            &tStatus{status: "SSL Handshake Failed", code: 525},
		InvalidSslCertificate:         &tStatus{status: "Invalid SSL Certificate", code: 526},
		RailgunError:                  &tStatus{status: "Railgun Error", code: 527},
		NoSuchUserHere:                &tStatus{status: "No Such User Here", code: 550},
		NetworkReadTimeoutError:       &tStatus{status: "Network Read Timeout Error", code: 598},
		NetworkConnectTimeoutError:    &tStatus{status: "Network Connect Timeout Error", code: 599},
		Decline:                       &tStatus{status: "Decline", code: 603},
	}
}