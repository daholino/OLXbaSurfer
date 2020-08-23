![OLXbaSurfer Logo](./resources/olxbasurfer-logo.png)

## OLXbaSurfer

`OLXbaSurfer` is a simple program that I wanted to write when I was searching for a car to buy. I always got annoyed when a listing pops up on OLX.ba and someone sees it before me and buys it.

This program will periodically search OLX.ba for given search query and it will send an email to you when new articles are added.

This program is using private OLX.ba API which can be changed in the future.

### Usage

To run this program you must specify some arguments when running it:

```
-clear
		If clear flag is set to true it will start the program with clean database.
-email string required
		Email where app will send notifies
-query string required
		Search query for OLX.ba
-smtp string required
		SMTP host URL
-smtp-pass string
		SMTP password
-smtp-port int
		SMTP port to connect to (default 587)
-smtp-user string
		SMTP username credential
-working-dir string
		Directory where database and log files are stored. (default "/var/OLXbaSurfer/")
-interval
		Search interval in hours (default 1)
```

Example of running it:

```
./OLXbaSurfer -query="apple tv tuzla" -clear=true -smtp="your SMTP host" -smtp-port=587 -smtp-user="your SMTP user" -smtp-pass="your SMTP pass" -email="your Email"
```

Some code is covered with tests, please keep in mind that this is the first time I am writing tests so I might improve them later when I gain more experience.

### Dependencies

- [Badger](https://github.com/dgraph-io/badger) - I used this neat key-value store library instead of regular SQL DB because this program only needs to check if something exists in database.
- [Cron](https://github.com/robfig/cron) - is used for scheduling searches for OLX.ba
