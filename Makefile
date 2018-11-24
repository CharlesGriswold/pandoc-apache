pandoc-apache: main.go
	go build

install: pandoc-apache
	install -o root -g root -m 644 files/pandoc-conf /etc/pandoc-apache.conf
	install -o root -g root -m 644 files/apache-conf /etc/apache2/conf-available/pandoc-apache.conf
	install -o root -g root -s pandoc-apache /usr/lib/cgi-bin/

dumpenv:
	install -o root -g root dumpenv.pl /usr/lib/cgi-bin/
