# Cuestionario

1. De que herramientas ó sitios se valdría para verificar la confiabilidad de un sitio web externo al cual
   debería conectarse?

- Lo primero en validar el si cuenta con un certificado ssl de seguridad, adicionalmente el metodo de autenticacion que
  utiliza para la seguridad como JWT, servidor Ouath
  o contraseñas simples, herramientas utilizadas para validar conozco:
  [url void](https://www.urlvoid.com/), [Sucuri](https://sitecheck.sucuri.net/) ,[Virus total](https://www.virustotal.com/gui/home/upload)

2.¿Qué Organismos de Internet velan por las buenas prácticas de desarrollo seguro? (nombre por lo
menos 2) Cuáles son estas prácticas?

- La organizacion mundial que se encarga de velar por las buenas prácticas de desarrollo es la
  OWASP fundations, [url OWASP](https://owasp.org/)
- Algunas de las buenas prácticas que sugieren en su documentation están:

* Codificacion segura: se debe validar las entradas y salidas de información del sistema, se debe validar que las
  librerias y paquetes
  de los terceros no esten deprecadas o con problemas de seguridad y siempre obtenerlas desde sitios oficiales.
* Fase de pruebas: Control de calidad de pruebas y seguridad, inspection del codigo por fases es decir pruebas unitarias
  y de integracion
  asi como pruebas de seguridad, en lo personal uso codesmell que ayuda a identicar falencias de codigo.
  Es importante también realizar pruebas que intenten bloquear el sistema como ataques de denegación de servicios
  inyeccion de codigo entre otros.

3. En otro proyecto nos solicitan exponer endpoints para integración de cuenta desde otras entidades
   financieras, en ese sentido si un cliente quiere vincular sus cuentas de BBVA desde otra entidad, que
   esquema de seguridad ofrecería? Cómo permitiría la integración sin poner en riesgo la seguridad del
   cliente? Cuáles serían los detalles de implementación del diseño a tener en cuenta? Detallar la
   necesidades de infraestructura y las distintas capas de seguridad a aplicar.

- Para la seguridad de los enpoint expuestos una buena práctica es el servidor Oauth que meneja la identification de los
  usuarios.
- Como arquitectura de seguridad de redes si no están en la nueve pude implementarse una DMZ para evitar ingreso no
  autorizados.
- Los detalles de la implementacion:

* Implementacion de servidor oauth.
* Implementacion de token dinamico para firmar las transaciones.
* Implementacion de cifrado llave publica privada con el cliente para encriptar la información y evitar salidas no
  autorizadas del sistema.
* Implementacion de observabilidad del las transaciones, paquetes, codigo y demás sistemas involucrados a fin de
  identificar por medio de tableros alertas, como perdida de datos, flujo inusuales a procesos.
* Implementacion de apache flink para identification de fraudes, esto funcionaria interno para ver los datos y
  estadistica en tiempo real.
* Implementacion de protocolos como GRPC para aumentar el nivel de seguridad entre servicios internos y de ser posible
  lo externos expuestos.

4. Dada la situación de fraude actual, se está poniendo muchísimo énfasis en la identificación lo más
   asertiva posible del cliente. Como muchos productos son gestionados totalmente online y entregados
   por empresas de envío , si se le pidiera un servicio de validación del cliente al que se le entrega el
   producto: que formas de verificación implementaría? qué servicios/orquestaciones expondría para los
   distintos canales? qué precauciones tendría en cuenta? qué requerimientos de plataforma/servicios
   solicitaría para sostener el servicio/orquestación? Como monitorizaría el servicio y que alarmas
   establecería ante detecciones de fraude?

- Para el caso de colombia, lo primero sería implementar una validation de identidad de huellas y cedula conectado a la
  registraduria nacional encargada de esta información.
- Adicionalmente, es impórtate la contravalidacion por medio de reconocimiento facial la cual se compara con la foto que
  está en la cedula del cliente.
- Si es posible me integro a servicios de validacion de identidad online como buro de credito o datacredito estas
  entidades por medio de codigo otp en movil y preguntas financieras validan si la persona
  es quien dice ser.
- Para monitorear el sistema puede ser una buena práctica implementar un tablero de logs y alertas de aws como
  CloudWatch, el cual permite ver el comportamiento de la aplicacion en fallos.
- Para monitorear tambien se puede implementar analiticas de eventos con firebase para poder ver la interaction en
  tiempo real del usuario con respecto a la aplicacion, esto se puede conectar
  con el apache kafka y flink lo que permite crear detention de fraudes, ejemplo la base de datos NEo4j que es de grafos
  permite el analysis de estos datos para ver si hay comportamientos atipicos
  como dos retiros de un mismo usuario en sitios diferentes, acceso en equipos no registrados, o intentos fallidos en el
  sistema, o para el caso de entregas o solicitudes en direcciones distintas a las
  registradas en las centrales de riesgo, o dane, entre otras muchas que se pueden sacar y volver un sistema de alertas
  capaz de identificar fraudes en tiempo real.
- También se puede implementar un sistema de notificaciones push o msn para informar de estas anomalias al usuario y
  evitar fraudes con información tar-dia.

5. ¿Con el siguiente supuesto que cierta data sensible debe ser intercambiada entre una aplicación
   mobile y su backend, debiendo asegurarse la confidencialidad en todo el camino de los datos: desde la
   app hasta servicios de middleware que originan la información (los cuales se encuentran detrás de la
   capa backend del canal), que mecanismos de seguridad sugeriría? Especificar y argumentar lo más
   detalladamente posible la respuesta.

- Para el caso expuesto ejemplo data de tarjetas de credito lo mejor es enviar la información encriptda como base de la
  solucion, es decir el cliente movil
  cuenta en su sistema una clave dinamica pública que es suministrada por el back por un servicio x, con esta key
  encripta los datos y los envia por el servicio rest pot, el back lee los datos
  y los desencripta segun la necesidad con su clave privada que es dinamica y unica por cada cliente movil, con esto me
  aseguro de que los datos no podran ser accedidos por personas no autorizadas,
  adicionalmente el servicio debe contar con authentication basada en oauth2 para evitar intentos no autorizados,
  es importante tambien si se puede utiliza un servicio como GRPC que por medio de los .protos asegura que la data no
  sea facil de leer por un humano.
- Para la parte de la infrastructura utilizar un dns inverso para evitar ataques asi como posibles puertos abiertos
  expuestos.

6. Ante el siguiente stack trace:
   ssl.AuthSSLX509TrustManager.checkServerTrusted:PKIX path building failed:
   sun.security.provider.certpath.SunCertPathBuilderException: unable to find valid certification path to requested
   target
   sun.security.validator.ValidatorException: PKIX path building failed:
   sun.security.provider.certpath.SunCertPathBuilderException:
   unable to find valid certification path to requested target
   at sun.security.validator.PKIXValidator.doBuild(PKIXValidator.java:323)
   at sun.security.validator.PKIXValidator.engineValidate(PKIXValidator.java:217)
   at sun.security.validator.Validator.validate(Validator.java:218)
   at com.sun.net.ssl.internal.ssl.X509TrustManagerImpl.validate(X509TrustManagerImpl.java:126)
   at com.sun.net.ssl.internal.ssl.X509TrustManagerImpl.checkServerTrusted(X509TrustManagerImpl.java:209)
   at ar.org.bbva.util.net.ssl.AuthSSLX509TrustManager.checkServerTrusted(AuthSSLX509TrustManager.java:96)
   at com.sun.net.ssl.internal.ssl.ClientHandshaker.serverCertificate(ClientHandshaker.java:1201)
   at com.sun.net.ssl.internal.ssl.ClientHandshaker.processMessage(ClientHandshaker.java:135)
   at com.sun.net.ssl.internal.ssl.Handshaker.processLoop(Handshaker.java:593)
   at com.sun.net.ssl.internal.ssl.Handshaker.process_record(Handshaker.java:529)
   at com.sun.net.ssl.internal.ssl.SSLSocketImpl.readRecord(SSLSocketImpl.java:943)
   at com.sun.net.ssl.internal.ssl.SSLSocketImpl.performInitialHandshake(SSLSocketImpl.java:1188)
   at com.sun.net.ssl.internal.ssl.SSLSocketImpl.writeRecord(SSLSocketImpl.java:654)
   at com.sun.net.ssl.internal.ssl.AppOutputStream.write(AppOutputStream.java:100)
   at java.io.BufferedOutputStream.flushBuffer(BufferedOutputStream.java:65)
   at java.io.BufferedOutputStream.flush(BufferedOutputStream.java:123)
   at org.apache.commons.httpclient.methods.EntityEnclosingMethod.writeRequestBody(EntityEnclosingMethod.java:506)
   at org.apache.commons.httpclient.HttpMethodBase.writeRequest(HttpMethodBase.java:2114)
   at org.apache.commons.httpclient.HttpMethodBase.execute(HttpMethodBase.java:1096)
   at org.apache.commons.httpclient.HttpMethodDirector.executeWithRetry(HttpMethodDirector.java:398)
   at org.apache.commons.httpclient.HttpMethodDirector.executeMethod(HttpMethodDirector.java:171)
   at org.apache.commons.httpclient.HttpClient.executeMethod(HttpClient.java:397)
   at org.apache.commons.httpclient.HttpClient.executeMethod(HttpClient.java:346)

   ¿Cómo lo solucionaría? Detallar paso a paso la resolución y que datos adicionales requeriría para
   resolverlo.

- Vaya a la URL en su navegador:
- Firefox: haga clic en la cadena de certificados HTTPS (el icono de candado justo al lado de la dirección URL). Haga
  clic en "more info" > "security" > "show certificate" > "details" > "export..". Elija el nombre y elija el tipo de
  archivo ejemplo.cer
- Chrome: haga clic en el icono del sitio a la izquierda de la dirección en la barra de direcciones, seleccione "
  Certificado" -> "Detalles" -> "Exportar" y guárdelo en formato "Certificado único, binario codificado en Der".
- Ahora tiene un archivo con el almacén de claves y debe agregarlo a su JVM. Determinar la ubicación de los archivos
  cacerts, p.

```sh
 C:\Program Files (x86)\Java\jre1.6.0_22\lib\security\cacerts.
 ```

- A continuación, importe el example.cerarchivo en cacerts en la línea de comandos (es posible que necesite el símbolo
  del sistema del administrador):

```sh
keytool -import -alias example -keystore  "C:\Program Files (x86)\Java\jre1.6.0_22\lib\security\cacerts" -file example.cer
```

- Se le pedirá una contraseña que por defecto eschangeit y reiniciar JVM/PC.

7. Dada una empresa que delega el desarrollo de ciertas funcionalidades/servicios a terceros, cómo
   implementaría el acceso de sus empleados sin comprometer los datos de autenticación de los mismos?

- Para este caso seria separar los ambientes de desarrollo y certificacion que sea con diferentes credenciales que las
  de produccion, adicionalmente un sistema de login con rol
  en el sistema para controlar quien ingresa y que realiza en el sistema, con bases de datos autenticadas.
- Adicionalmente, si es posible en el sistema de gestion de archivos, correo, y canales de comunicacion utilizar doble
  autenticacion con el movil o sistemas online para evitar ingresos no
  autorizados a sistemas criticos para una empresa.

