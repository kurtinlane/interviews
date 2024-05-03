### Next steps 

Simplify all the copy pasted functions into 1. Where a URL and a object get passed in and the object gets populated with the data from the URL 

Don't love the mapping from starwars object to our object, but most apis aren't great and things change without anyone saying anything. This also allows us to use different APIs in the future. Also our contract doesn't break if the starwars api changes (cuz it will).

Would want a cleaner way to shutdown the HTTP server and close out any current requests that are being handled and close out any DB connections etc gracefully 

WOuld move them into their appropriate packages (to match whatever the status quo is, don't really care about the structure so long as its consistent from repo to repo)

Would also fix some of the names, so the next poor sucker who has to work on this can just read the code to see whats going on. 

IFFFFF anyone actually ends up using this, then I'd put some metrics into it and see what is slowest and then use some channels or go funcs to speed this bad boi up. But for the extra complexity right now it ain't worth it unless this is getting used.

