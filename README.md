Postqres Ayarları
--
*Docker kurulum*

    docker run --name tiny-url-api-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres


*DB Bilgileri*
 - DB_USERNAME = postgres 
 - DB_PASSWORD = postgres
 - DB_HOST = localhost
 - DB_PORT = 5432 
 - DB_TABLE = postgres
                                                
                                            

Layouts

Cmd/Server
-
Projenin ana uygulamalarını içerir. İçerisinde main.go program başlangıç paketimiz bulunmaktadır. Başlangıçta Veritabanı, Migration ve Handler ve Endpointler ayağa kaldırılmaktadır.

Internal 
-
Özel uygulama ve kütüphane kodunu içerir
	

 - **database/database.go**
		 - Veritaban bağlantısını ve db instance veren paket.
 - **database/migration.go**
		 - Veritabanı modellerinin migration işleminin yapıldığı paket.
 - **tinyurl/tinyurl.go**
		 - Url api ve veritabanı arasında CRUD işlemlerini yapan paket.
 - **transport/http/handler.go**
		 - Api Endpointlerinin, Middleware, Response gibi servislerin olduğu paket.
 - **transport/http/jwt.go**
		 - JWT token oluşturma ve validasyon işlemlerinin olduğu paket.
 - **transport/http/middleware.go**
		 - Birden fazla middleware yazabilecek paket. İçerisinde örnek olsun diye Log middleware eklenildi.
 - **transport/http/tinyurl.go**
		 - Url endpointlerinin bulunduğu paket.
 - **transport/http/user.go**
		 - User endpointlerinin bulunduğu paket.
 - **user/user.go**
		 - User  api ve veritabanı arasında kullanıcı oluşturma ve kullanıcı sorgulama işlemlerini yapan paket.
 - **utils/random.go**
		 - Random herhangi bir fonksiyonları yazılan paket. İçerisinde Random URL oluşturan fonskiyon bulunmaktadır.


