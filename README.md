# Notelet Server
A simple note manegemnt system for creating, storing, viewing and searching useful notes. 
The service provides REST API aimed at managing the notes and a template suitable for SPA (single-page application) client.  
It is written in Go programming language.
## Routes
Notelet server REST service can be acessed:
* > '<domain_name>/api

SPA template:
* > <domain_name>/

## REST Api
REST Api provides the following services:
| HTTP verb     | Path          | Service |
| ------------- |:-------------:| -----:|
| GET    | /notes<?filter=searchstring> | Retrieve the list of notes optionallly filtered by searchstring |
| GET    | /notes/{id} | Retrieve a note with the given id |
| POST   | /notes      | Add a new note |
| DELETE | /notes/{id} | Remove a note with the given id
| PUT | /notes/{id} | Replace a note with the given id