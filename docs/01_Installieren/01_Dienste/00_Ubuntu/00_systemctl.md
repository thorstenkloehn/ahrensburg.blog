

## Allgemeine Aufbau Datei

```
[Unit]
# Abschnitt wird im Artikel systemd/Units beschrieben

[Service]
Type=simple
ExecStart=/PFAD/ZUM/BEFEHL/befehl

[Install]
WantedBy=multi-user.target

```
## Optionen für die [Service]-Sektion

* ExecStart    - 	der Befehl, der beim Start der Unit ausgeführt wird
* ExecStartPre - 	der Befehl, der vor dem Start der Unit (also vor der Ausführung von ExecStart) ausgeführt werden soll
* ExecStartPost - 	der Befehl, der nach dem Start der Unit (also nach der Ausführung von ExecStart) ausgeführt werden soll
* WorkingDirectory - 	legt das Arbeitsverzeichnis fest, in dem die Prozesse ausgeführt werden. Das Verzeichnis muss als absoluter Pfad angegeben werden oder als ~. Bei letzterem wird das Homeverzeichnis des im Schlüssel User angegebenen Nutzers gewählt.
* User -	legt fest, unter welchem Benutzer der Service laufen soll (Standard: root)
* Group -	legt fest, unter welcher Gruppe der Service laufen soll

## Beispiel 

```

[Unit]
Description=ahrensburg.digital

[Service]
WorkingDirectory=/Server/ahrensburg.digital
ExecStart=/Server/ahrensburg.digital/ahrensburg.digital
Restart=always
# ahrensburg-dital
RestartSec=10
KillSignal=SIGINT
SyslogIdentifier=ahrensburg.digital
User=root

[Install]
WantedBy=multi-user.target


```
## Dienst Einstellung

* sudo cp -u ahrensburg.service /etc/systemd/system/ahrensburg.service
* nano /etc/systemd/system/ahrensburg.service

## Server Dienst Aktivieren
```
sudo  systemctl enable ahrensburg.service


```


## Komplett Einleitung

* [Komplett Einleitung](https://www.freedesktop.org/software/systemd/man/systemctl.html)

