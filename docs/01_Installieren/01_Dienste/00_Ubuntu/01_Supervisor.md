
## Installieren 
```
sudo apt-get install supervisor

```

## Beispiel Code

```
[program:python] ;Programm Name
directory=/Server/flask ;Pfad vor der Ausführung wechseln
command=/usr/bin/gunicorn app:app -b localhost:6000 ;Der Pfad des überwachten Prozesses
numprocs = 1; Prozess starten
autostart=true ; Beginnen Sie mit dem Start von Supervisord
autorestart=true ;  automatischer Neustart
startretries = 10; Die maximale Anzahl von Wiederholungsversuchen, wenn der Start fehlschlägt
Exitcodes = 0; normale Exitcodes
Stoppsignal = KILL; Das Signal, mit dem der Prozess abgebrochen wird
stopwaitsecs = 10; Wartezeit vor dem Senden von SIGKILL
```

## Erstellen Supervisor

```

nano etc/supervisor/conf.d/python.conf

```

## Starten Supervisor

* sudo supervisorctl reread
* sudo service supervisor restart

