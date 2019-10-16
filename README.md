# card_collection

## lookup

Basic HTTP service that provides two HTTP GET endpoints:

* /:cardname (eg. /opt or /ajani+adversary+of+tyrants)
* /:setcode/:collector_number (eg. /dom/60 or /m19/3)

Takes the variables and looks them up via the Scryfall API