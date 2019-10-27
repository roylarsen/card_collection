# card_collection

## lookup

Basic HTTP service that provides two HTTP GET endpoints:

* /:cardname (eg. /opt or /ajani+adversary+of+tyrants)
* /:setcode/:collector_number (eg. /dom/60 or /m19/3)

Takes the variables and looks them up via the Scryfall API

## user mgmt

Service for handling user management functions.

Provides two POST endpoints:
* /user_add
* /login

#### Notes:

As I'm learning this stuff, I'm borrowing a lot of code from https://echo.labstack.com/cookbook/twitter for my user management stuff.
Doing user management stuff and JWT is new for me, so, I'm borrowing parts of that code and then modifying it (partly to force me to learn it, but also to use the new(-ish) mongodb driver).