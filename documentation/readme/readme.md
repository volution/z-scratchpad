



![logo](./documentation/logo.png)




--------------------------------------------------------------------------------




# `z-scratchpad` -- lightweight Go-based notes tool


> ## Table of contents
>
> * [about](#about); [status](#status); [screenshots](#screenshots);
> * [documentation](#documentation);  [install](#install);
> * [features (and anti-features)](#features); [performance](#performance);
> * [how? (concepts, and inner workings)](#how);
> * [why? (history, and reasons)](#why);
> * [contributions](#contributions); [licensing](#license);




--------------------------------------------------------------------------------




## <span id="about">About</span>

`z-scratchpad`, as the title says, is a lightweight and highly customizable tool
implemented in pure Go (thus portable to most POSIX compliant OS's)
that allows one to easily take notes,
from the unorganized post-it pile, reusable copy-paste snippets,
memos, document drafts,
up to the highly organized (and carefully tagged and categorized) project specific documents.

Most importantly, `z-scratchpad` does not use a database (neither embedded nor client-server);
instead it uses plain text files with minimal syntax, on a plain file-system,
with minimal requirements on file and folder organization.
Thus, the actual documents can be easily searched and edited with tools like `nano`, `grep`, `sed`, etc.,
versioned with Git, Mercurial, etc.,
synchronized with `rsync`, Dropbox, etc.,
and interacted with like any other plain text file.

It provides a CLI (command line), TUI (terminal interface, a.k.a. "curses" interface), and a GUI.
However, the TUI / GUI is delegated to one's preferred tools
(like `nano`, `vim`, `sublime_text`, etc., for editing, and `fzf`, `dmenu`, `rofi`, etc., for menus).

It also provides a simple WUI (web interface via a built-in HTTP server) for browsing and viewing.
Although the WUI is not intended to be the primary interface, and was not designed to be exposed to the network.
However, for publishing on the internet it provides a simple HTML static site export.

It is tailored to individual use, although, given it uses only plain text files, one could leverage Git or any other versioning system.
Also, it does not require running a daemon or server in the background.

One can think of `z-scratchpad` as the merger of
[Notational Velocity](https://notational.net/)
(or [nvALT](https://brettterpstra.com/projects/nvalt/))
with [MoinMoin](https://moinmo.in/) (or other wiki's),
that uses one's preferred editor and UI tools.

What `z-scrpatchpad` is not:
* a fancy mind-mapping, org-mode, or the latest buzzword note taking tool;
* a fully-fledged wiki;  it does have a web interface, but that is tailored for browsing and viewing;  (and it should never be exposed to the network;)
* a document management system;

For more features, and anti-features, please see the [dedicated section](#features).

Also see the following useful sections:
* [status](#status) -- the current implementation status;
* [install](#install) -- how to download and install it;
* [documentation](#documentation) -- how to use and configure it;
* [performance](#performance) -- what means to be "fast-enough";
* [UI look-and-feel](#ui) -- how the UI works, and what can one expect when interacting with it;
* [how it works](#how) -- the various use-cases it covers, and how it works internally;
* [why this tool](#why) -- describes how I've reached to implement this tool, and why it behaves like it does;

Finally, it is open-source, licensed under GPL 3 or later.
Please see the [contributions](#contributions) and [licensing](#license) sections.




--------------------------------------------------------------------------------




## <span id="status">Status</span>

> **WIP** (work in progress)

At the moment `z-scratchpad` is still under heavy development.

That being said, I'm using it for all my note taking,
from personal notes, to for-work project specific documents.

Moreover, given that it just provides the glue between one's favorite UI tools (editors and such),
if it breaks nothing is lost.
All one's files are stored on the file-system, in plain text files,
thus usable with generic file-management tools.




--------------------------------------------------------------------------------




## <span id="screenshots">Screenshots</span>


### WUI in Firefox

The WUI showing this readme rendered in Firefox browser (under Xorg):

![wui-firefox-document](./documentation/screenshots/wui-firefox-document.png)


### WUI under terminal

The WUI showing this readme rendered in [links](http://links.twibright.com/) (under URxvt terminal) (scrolled 1 page to skip the document header):

![wui-links-document](./documentation/screenshots/wui-links-document.png)


### Menus under Xorg

Selecting a document to edit with [`rofi`](https://github.com/davatorium/rofi) (under Xorg):

![select-rofi-document](./documentation/screenshots/select-rofi-document.png)


### Menus under terminal

Selecting a document to edit with [`fzf`](https://github.com/junegunn/fzf) (under URxvt terminal with [`tmux`](https://github.com/tmux/tmux) support):

![select-fzf-document](./documentation/screenshots/select-fzf-document.png)




--------------------------------------------------------------------------------




## <span id="documentation">Documentation</span>

> **WIP** (work in progress)

Besides what is available by running `z-scratchpad help` there is no other documentation at the moment.

That being said, just run the following and start experimenting with the commands.
(If there is need for documentation, besides the frugally `-h` for each command, I have failed in one of the mandatory requirements, that of being "simple to use".)

For how to download and install it see the [dedicated section](#install).

Get some help:
~~~~
z-scratchpad -h
z-scratchpad create -h
z-scratchpad edit -h
z-scratchpad list -h
~~~~

Initialize the notes store:
~~~~
# create an empty folder for the notes store
mkdir ./some-notes

# mark the folder as the notes store
touch ./some-notes/.z-scratchpad

# switch to the notes store
# (alternatively, add `-C ./some-notes` to all commands)
cd ./some-notes
~~~~

Create a note (with a random identifier prefixed by the current date):
~~~~
z-scratchpad create
~~~~

Create a note with a custom identifier:
~~~~
z-scratchpad create -d some-identifier
~~~~

The note syntax is very simple:
* use the first line, prefixed with `##`, as a document title;  (e.g. `## some title`;)
* leave at least one empty line;  (to separate the note header from the body;)
* write anything afterwards;
~~~~
## some title

some content
~~~~

Edit a note by selecting it from a menu:
~~~~
z-scratchpad edit -s
~~~~

Edit a note by using its identifier:
~~~~
z-scratchpad edit -d some-identifier
~~~~

List all notes (either identifiers, titles, or paths):
~~~~
z-scratchpad list
z-scratchpad list -w identifier
z-scratchpad list -w title
z-scratchpad list -w path
~~~~

List all the files in the notes store:
~~~~
ls -a -1
~~~~
~~~~
2021-08-13--6d2ace99.txt
some-identifier.txt
.z-scratchpad
~~~~

Start the WUI server:
~~~~
z-scratchpad server
~~~~

Open a note in the browser by selecting it from a menu:
~~~~
z-scratchpad browse -s
~~~~

Open a note in the browser by using its identifier:
~~~~
z-scratchpad browse -d some-identifier
~~~~




--------------------------------------------------------------------------------




## <span id="install">Installation</span>

As mentioned many times, `z-scratchpad` is a single self-contained executable,
thus it can be deployed by downloading the executable and placing it somewhere on the `$PATH`.

Traditionally one should place it inside the `/usr/local/bin` folder (available on many POSIX compliant OS's).
Alternatively, one can place it in one's `$HOME/bin` folder and add that to `$PATH`.

### Releases

The self-contained executable is available from the `z-scratchpad` GitHub repository [releases](https://github.com/volution/z-scratchpad/releases):
* [v0.0.1](https://github.com/volution/z-scratchpad/releases/tag/v0.0.1) -- the first preliminary preview release:
  * [v0.0.1 for Linux](https://github.com/volution/z-scratchpad/releases/download/v0.0.1/z-scratchpad--linux--v0.0.1) -- tested on OpenSUSE, and used by me in my everyday work;
  * [v0.0.1 for OSX](https://github.com/volution/z-scratchpad/releases/download/v0.0.1/z-scratchpad--darwin--v0.0.1) -- untested;
  * [v0.0.1 for FreeBSD](https://github.com/volution/z-scratchpad/releases/download/v0.0.1/z-scratchpad--freebsd--v0.0.1) -- untested;
  * [v0.0.1 for OpenBSD](https://github.com/volution/z-scratchpad/releases/download/v0.0.1/z-scratchpad--openbsd--v0.0.1) -- untested;

Also, each of these files are signed with my PGP key `5A974037A6FD8839`, thus do check the signature.

### Download and verify

The following is an example how one could download, verify, and deploy `z-scratchpad`:

* import my PGP key:
~~~~
curl -s https://github.com/cipriancraciun.gpg | gpg2 --import
~~~~
~~~~
gpg: key 5A974037A6FD8839: public key "Ciprian Dorin Craciun <ciprian@volution.ro>" imported
gpg: Total number processed: 1
gpg:               imported: 1
~~~~

* download the executable and signature (replace the `linux` token with `darwin` (for OSX), `freebsd` or `openbsd`):
~~~~
curl -s -L -S -f -o ./z-scratchpad \
    https://github.com/volution/z-scratchpad/releases/download/v0.0.1/z-scratchpad--linux--v0.0.1

curl -s -L -S -f -o ./z-scratchpad.asc \
    https://github.com/volution/z-scratchpad/releases/download/v0.0.1/z-scratchpad--linux--v0.0.1.asc
~~~~

* verify the executable:
~~~~
gpg2 --verify ./z-scratchpad.asc ./z-scratchpad
~~~~

* **check that the key is `58FC2194FCC2478399CB220C5A974037A6FD8839`**:
~~~~
gpg: assuming signed data in './z-scratchpad'
gpg: Signature made Sat Aug 14 17:12:39 2021 EEST
gpg:                using DSA key 58FC2194FCC2478399CB220C5A974037A6FD8839
gpg: Good signature from "Ciprian Dorin Craciun <ciprian@volution.ro>" [unknown]
gpg:                 aka "Ciprian Dorin Craciun <ciprian.craciun@gmail.com>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 58FC 2194 FCC2 4783 99CB  220C 5A97 4037 A6FD 8839
~~~~

* change the executable permissions:
~~~~
chmod 0755 ./z-scratchpad
~~~~

* copy the executable on the `$PATH`:
~~~~
sudo cp ./z-scratchpad /usr/local/bin/z-scratchpad
~~~~

* check that it works:
~~~~
z-scratchpad --version
~~~~
~~~~
* version       : 0.0.1
* executable    : z-scratchpad
* build target  : release, linux-amd64, go1.16.7, gc
* build number  : 2025, 2021-08-13-11-56-08
* code & issues : https://github.com/volution/z-scratchpad
* sources git   : 1abeee1c76fc4a40e1465e4810be0258992d1815
* sources hash  : 51382d0da05c3e129fc73eac59754fad
* uname node    : some-workstation
* uname system  : Linux, 5.some-version, x86_64
~~~~

### Build from source

Alternatively, one can just build the executable themselves (for example to get the latest unreleased changes).

* clone the repository:
~~~~
git clone https://github.com/volution/z-scratchpad ./z-scratchpad
~~~~

* switch to the sources folder:
~~~~
cd ./z-scratchpad/sources
~~~~

* build the executable:
~~~~
go build -o ./z-scratchpad ./cmd/z-scratchpad.go
~~~~

* copy the executable on the `$PATH`:
~~~~
sudo cp ./z-scratchpad /usr/local/bin/z-scratchpad
~~~~

* check that it works:
~~~~
z-scratchpad --version
~~~~
~~~~
* version       : 0.0.1
* executable    : z-scratchpad
* build target  : release, linux-amd64, go1.16.7, gc
* build number  : 2025, 2021-08-13-11-56-08
* code & issues : https://github.com/volution/z-scratchpad
* sources git   : 1abeee1c76fc4a40e1465e4810be0258992d1815
* sources hash  : 51382d0da05c3e129fc73eac59754fad
* uname node    : some-workstation
* uname system  : Linux, 5.some-version, x86_64
~~~~




--------------------------------------------------------------------------------




## <span id="features">Features (and anti-features)</span>

### Features and requirements

The following are the main requirements, sorted by priority, that I have in mind while implementing or extending `z-scratchpad`:

* (**mandatory**) stores the documents in plain text files, with minimal requirements on the file-system structure;  (if tomorrow `z-scratchpad` disappears, no one should lose anything, and should be able to easily migrate to other tools;)
* simple to use;  (once configured, although it can be easily used even without a configuration file, based on sensible defaults;)
* simple to integrate with one's favorite tools;  (from preferred editor to preferred browser;)
* (**mandatory**) works under the terminal (i.e. TUI);
* (**mandatory**) works under Xorg / Wayland (i.e. GUI);
* (**mandatory**) does not implement any TUI / GUI;  (instead delegates everything to other tools like `fzf`, `dmenu` or `rofi`;)
* (**mandatory**) provides a CLI;  (thus can be integrated in custom workflows, for example `bash` scripts and other scripting tools;)
* (**mandatory**) does not require running a daemon or server;
* provides an HTTP interface accessible from a browser  (i.e. WUI);
* provides an HTML static site export;  (for selected documents, thus one should be able to mix private and public documents;)
* support for [CommonMark](https://commonmark.org/);  (that is of importance mainly for the WUI and the HTML export;)
* (**mandatory**) portability to many of the POSIX compliant OS's, especially Linux (my main environment), OSX, OpenBSD and FreeBSD;
* (**mandatory**) single executable, place it anywhere deployments;  (anything that is required to have it running, including assets for the WUI, should be embedded in the binary;)

Note that some requirements are marked with "mandatory" although are not at the top.
The reason is that although during implementation they should be maintained, compromises can be made (for example in terms of performance).
On the other hand, the higher some are on the list, the fewer compromises should be made.


### <span id="anti-features">Anti-features</span>

Conversely, there are also some negative requirements, or anti-features, that I keep in mind:

* (**mandatory**) does not support any non-text documents;  (it doesn't care what is inside the document, i.e. its syntax, as long as it's a plain text file;)
* (**mandatory**) does not support any non-ASCII or non-UTF-8 documents;
* (**mandatory**) does not implement any TUI / GUI;
* (**mandatory**) does not implement any form of encryption;  (always use full disk encryption, always use encrypted swap, always use memory backed temporary folders;  always use tools that focus just on cryptography (like [GnuPG](https://gnupg.org/));  never use fancy tools that provide "encryption" features;)
* (**mandatory**) does not support multiple users;  (although one could use different instances;)
* does not provide support for attachments;  (one should use other means to store files, and link them;)
* does not provide extensions to the CommonMark syntax (except perhaps the quasi-standard ones introduced by GitHub, i.e. [GFM](https://github.github.com/gfm/), and supported by many current parsers);
* does not prioritize support for alternative syntaxes to CommonMark (or the no-syntax plain text alternative currently implemented);
* does not prioritize exposing the HTTP interface to the network;  (it should always listen on `localhost`;  if one needs to expose it to the network, please use a reverse proxy (like [HAProxy](https://www.haproxy.org/));)
* does not prioritize editing (and other workflows) via the WUI;
* does not prioritize support for images (or other media), especially in the WUI and HTML export;
* does not prioritize Windows support;  (in theory it should work, however the Windows ecosystem lacks many of the tools relied-on by the TUI / GUI;)
* does not prioritize built-in advanced workflows;  (however, by using the CLI interface one can implement in his favorite scripting language any workflow one desires;)


### <span id="ui">UI considerations</span>

The careful reader might see that I've listed "does not implement any TUI / GUI" twice, both in features and anti-features, it was not a mistake.
`z-scratchpad` should limit its UI requirements to the following primitive operations that can be provided by external tools.

* editing a plain text file;  (using one's preferred editor, from `nano`, `vim`, and `emacs` to `howl`, `sublime_text`, and `vscode`;)
* viewing a plain text file;  (which can be easily solved by the same editor as above;)
* selecting an option from a non-hierarchical menu;  (using for example `fzf` under the terminal, or `dmenu` or `rofi` under Xorg;)
* viewing an HTML file if one uses the WUI;  (the browser should not be used for any other purposes as part of other workflows;)

Anything that is not on this list should be made to fit a workflow based on these primitives.




--------------------------------------------------------------------------------




## <span id="performance">Performance</span>

Given that `z-scratchpad` should run without a server, and furthermore that the actual notes are stored on the file-system,
there is quite a lot to be asked in terms of performance...

So I would start to describe some acceptable goals and limitations:
* it is expected that one doesn't have more than a couple of thousand notes in the same instance;  (for performance profiling I use ~15K documents, all in one folder;)
* it is expected that one uses SSD's, and that the system isn't under such heavy memory pressure that the buffer cache is starving;
* it is expected that write operations (like creating or changing a document) happen less frequently than read operations (like listing, browsing, or just viewing documents);
* it is expected to trade memory (both RAM or disk) in favor of CPU time;  (although it should be within reasonable bounds;)
* the implemented caching mechanism should be simple, as to keep the overall code simple and bug-free;  (as one once said, "there are two hard things in computer science, cache invalidation and naming things";)
* cache invalidation requires to walk the file-system;  (we can't use the kernel's file-system notification mechanism, mainly because we don't have a server running in the background;)
* if one changes (adds, removes, or edits) the notes files without using the tool, one should manually call the tool to invalidate the cache;
* also let's take [Jakob Nielsen's advice](https://www.nngroup.com/articles/powers-of-10-time-scales-in-ux/) with regard to delays in user experience:
  * 0.1 seconds (100 milliseconds) -- "users feel like their actions are directly causing something to happen on the screen";
  * 1 second -- "users feel like the computer is causing the result;  although they notice the short delay, they stay focused on their current train of thought";
  * 10 seconds -- "it breaks the user's flow";
* at the moment, the tool requires all libraries, documents and index to be present in Go's memory as objects;  (there is no lazy loading;)

That being said, in my initial performance profiling, with a library composed of ~15K CommonMark documents (generated via [lorem-markdownum](https://github.com/jaspervdj/lorem-markdownum)),
I've obtained the following times:
* walking the file-system, parsing the notes meta-data (not the CommonMark syntax), and indexing the entire library takes ~0.7 seconds (i.e. 700 milliseconds);
  * a good chunk of this time is spent computing document fingerprints (based on cryptographic hashes);
  * another good chunk is spent actually interacting with the file-system;  (thus can't be optimized away;)
* loading an already cached index takes ~0.05 seconds (i.e. 50 milliseconds);
  * the majority of the time is spent in deserializing the index cache file;

In order to obtain these numbers I had to resort to quite a few optimizations:
* replaced `SHA1` (and before that `SHA256`) with [`blake3`](https://github.com/zeebo/blake3) with AVX2 and SSE4.1 acceleration;
* disabled the Go garbage collector while we are parsing and indexing the notes;  (not much garbage is actually generated;)
* used memory-mapped files (for the serialized index cache file);
* used polled `byte.Buffer` instances, grouped by sizes;
* quite a few unsafe operations so that temporary `[N]byte` buffers don't escape the stack, and thus don't incur allocations;

Therefore, at the moment, without complicating the code too much, this is the best one can achieve.
(There is one low-hanging fruit, that of parallelizing the document parsing, however that only reduces the latency, not the overall CPU usage.)

However, even with the current state, I think that the performance objective was achieved.




--------------------------------------------------------------------------------




## <span id="how">How, concepts, and inner workings...</span>

> **WIP** (work in progress)

In this section I mainly describe what use-cases `z-scratchpad` should cover,
how it should integrate in one's environment,
and how it should store one's documents.

If one is interested in why I've reached this model, please see the next section on ["why"](#why).


### Concepts

`z-scratchpad` uses the following concepts:
* **instance** -- mainly identified with a single configuration file;  (one can have many instances, each using a different configuration file;)
* **document** -- an individual plain text UTF-8 file, composed of lines, that has a header and a body;
  * **document header** -- the first contiguous block of non-empty lines, having a simple syntax, used to give the document a title and some other meta-data;
  * **document body** -- all the other lines following the header, separated by at least one empty line;  the syntax of the body depends on the document format;
  * **document format** -- how should the body be parsed (mainly when exporting to HTML);  currently there are three supported formats:  CommonMark, snippets, and just text;
  * **document title** -- one or more "titles" that are used mainly in the UI to select a document;  (multiple documents can have the same title, although it is not advisable;)
  * **document identifier** -- a token (with strict syntax) that uniquely identifies a document within its library;
  * **document snapshot** -- an optional backup of a document, created just before editing it;
* **library** -- a set of documents;  one can have multiple libraries inside the same instance;
  * **library paths** -- one or more folders that are recursively walked to identify documents;
  * **library create path** -- exactly one (or none) folder where new documents should be placed in;  (think of this like the `inbox` folder;)
  * **library identifier** -- a token (with strict syntax) that uniquely identifies a library within an instance;
  * **library configuration** -- various properties that change certain behaviors when interacting with this library;
  * **default create library** -- each instance can have exactly one library where new documents are placed by default (if no library is specified);  (think of this like the `inbox` library;)
* **menus** -- each instance can have configured one or more flat menus;
  * **menu item** -- each entry in a menu, with a display label, and a command with arguments;
  * **menu command** -- a sub-set of the available commands that can be called from within the menu;  (mainly creating, editing, searching and opening documents, plus showing other sub-menus;)
  * nested menus -- although it does not support "nested" or "hierarchical" menus, one can call another menu as a command;  (thus one can implement arbitrary menu paths;)


### Use-cases and workflows

Creating a new document:
* `z-scratchpad create` -- if there is configured a default create library, an new document with a random name (prefixed with the current date) is created under that library's create path, and the preferred editor is opened with the corresponding file;
* `z-scratchpad create -l some-library` -- as above the document would be created in the given library's create path;
* `z-scratchpad create -s` -- the user is prompted to select a library where a document should be created;
* `z-scratchpad create -l some-library -d some-document` -- a document with the given identifier is created in the given library;

Editing an existing document:
* `z-scratchpad edit -s` -- the user is prompted to select a document from all available libraries, and then the preferred editor is opened with the corresponding path;
* `z-scratchpad edit -l some-library -s` -- as above, but the document selection is limited to the given library;
* `z-scratchpad edit -l some-library -d some-document` -- the document with the given identifier and in the given library is opened for edit;

Opening an existing document in the preferred browser:
* just replace `edit` with `browse`;

Integrating in other scripts:
* `z-scratchpad list -t library` -- lists the identifiers of available libraries;
* `z-scratchpad list -t library -f json` -- the same as above, but output a JSON array;
* `z-scratchpad list -t library -l some-library -w path` -- list all the store paths for the given library;
* `z-scratchpad list -t document -l some-library -w path` -- list all the document paths for all the documents in the given library;
* `z-scratchpad grep -t some-token-a -t some-token-b -W body -w path` -- list all the document paths, whose title contain any of the given tokens;
* `z-scratchpad grep -t some-token-a -t some-token-b -W body -w path` -- list all the document paths, whose bodies contain any of the given tokens;
* `z-scratchpad export -l some-library -d some-document -f source` -- export the given document's source code;  (with the document header canonicalized;)
* `z-scratchpad export -l some-library -d some-document -f html` -- export the given document's body rendered as HTML (only the actual body, that could be included in for example `<main>...</main>`);


### TUI vs GUI

`z-scratchpad` tries to detect if it is running under a terminal or Xorg:
* it considers running under a terminal if all these conditions are met:
  * the `TERM` environment variable is set (and not equal with `dumb`);
  * the `stderr` file descriptor is a TTY;
  * terminal access is not disabled (for example via configuration, or running as a server;)
* it considers running under Xorg if all these conditions are met:
  * it does not consider running under a terminal;  (i.e. TUI has precedence over GUI;)
  * the `DISPLAY` environment variable is set;
  * Xorg access is not disable (for example via configuration;)

Depending on whether it considers running under a terminal or Xorg, it tries to use different tools (for editing, selecting, etc.)

However one can always set the same tools for both terminal or Xorg configuration properties.


### WUI HTML usability

The WUI HTML was designed as simple as possible for two reasons:
* to keep the amount of clutter (links, details, etc.) as low as possible;
* to have it usable even without CSS, **especially in a terminal browser**;  (it works best with `links`, `w3m` and `lynx`;)
* (hopefully, screen readers and other assistive technologies are able to work with it without a problem;)

As stated multiple times throughout this document, `z-scratchpad` does not intend to be a "publishing wiki",
thus the interface doesn't need to do much except render the documents for easier viewing.


> **TBC** (to be continued)



--------------------------------------------------------------------------------




## <span id="why">Why, and history and reasons...</span>

I've been a long time user of the [MoinMoin](https://moinmo.in/) wiki,
my earliest document being from April 2008,
and I've used it for anything, from writing research papers,
to `bash` snippets, to-do lists,
and even as a store for encrypted credentials.
However, somewhere in 2015 I've started moving away from it to something else.

The main issue I had with MoinMoin was the editing experience.
Not the syntax, because I liked the MoinMoin [one-of-a-kind syntax](https://moinmo.in/HelpOnMoinWikiSyntax),
and besides that it also had support for many others
(a favorite of mine being [reStructuredText](https://en.wikipedia.org/wiki/ReStructuredText)).
When I say the editing experience I mean actually editing the wiki markup,
which required me to use the provided `<textarea>`,
that not only was incredibly small
(by default it used 20 rows with 80 columns, thus covering perhaps at most 50% of the display of even a small laptop),
but also gave me the experience of the Windows 95 Notepad...

But I've managed by using a Firefox plugin called [It's All Text!](https://github.com/docwhat/itsalltext),
that unfortunately somewhere in 2017 with the release of Firefox 57 stopped working.
This plugin added a small icon to one of the `<textarea>` corners (for any site)
that when pressed would open my favorite plain text editor,
and then listened for saves and refreshed the `<textarea>` contents.

However, even with this small improvement, it all just didn't work with my workflow.
I have (and had) the habit of writing (and saving) anything, from one-time `bash` snippets to personal notes on a phone-call.
Basically, I've used my wiki as both a "highly structured document editing and publishing platform",
but also as a "blackboard full of post-its";  it worked for the former, id really didn't for the latter.

So around January 2015 I'we wrote a simple `bash` script, called `x-scratchpad` that would do the following three simple things:
* call it with `create` (or no arguments),
  and it would generate a random 8 hex-character identifier,
  create a text file (in a certain hard-coded "store" folder),
  then open the preferred editor with that file;
* call it with `open`,
  and it would list all the file names in the store,
  concatenate that with the first line in each file,
  pipe that through `dmenu` to allow the user to select a file (by title),
  then open the same preferred editor with that file;
* call it with `search`,
  and would do the same thing as `open`, but instead of using only the first line,
  it would use all (non-empty) lines in each file;
  (thus a basic "full-text-search";)

Granted, this is not a "new and unique" concept, OSX users had [Notational Velocity](https://notational.net/) and [nvALT](https://brettterpstra.com/projects/nvalt/),
however Linux users only had [Zim](https://zim-wiki.org/)...

This small `bash` script allowed me to cover the "blackboard full of post-its" use-case.
However, using it was such a breeze as compared to the MoinMoin workflow (especially creating new documents),
that very soon I've started using it instead of my wiki.
(By carefully using titles such as `project / topic / document`, I could have both use-cases in the same folder.)
At the moment I have around 2K documents (and I guess in the last 5 years perhaps I've cleaned at least as many),
and I use it for anything.

Since 2015 until now (that is 2021), I've made only minor improvements,
and although I've been a happy user, it started to show it's pain points:
* first, it got slow;
  read the first line in each file, prefix that with the file name,
  and do this for 2K documents takes around a second or two
  (even on an hot-cache SSD);
* secondly, it blocked me from publishing some of my "unorganized" notes and snippets;
  (something that MoinMoin excelled at;)

I did investigate lots of alternatives, but all had their issues which I would summarize as:
* focused on publishing full-blown wikis, with the same trade-offs as MoinMoin;
  (I need not add that perhaps 99% of these were written in NodeJS...)
* focused on local note taking, with lots of bells-and-whistles,
  but lacking integration with custom plain text editor;
  (not to mention that perhaps 99% of these were using Electron, thus written in NodeJS...)
* focused exclusively on terminal (i.e. CLI) interaction;

I did encounter a few tools that had some of my desired traits,
but none had everything "just right":
* [neuron](https://github.com/srid/neuron), and perhaps its rewrite [emanote](https://github.com/srid/emanote),
  that is the closest to the tool I've written;
* [nuttyartist/notes](https://github.com/nuttyartist/notes), GUI, similar with nvALT or Notational;
* [zk](https://github.com/mickael-menu/zk) and [pimterry/notes](https://github.com/pimterry/notes), focused on terminal and CLI;

(If one would look carefully at these projects,
one would observe that all are written in **compiled languages**,
that yield actual **native executables**,
and none of them are written in NodeJS...) :)




--------------------------------------------------------------------------------




## <span id="contributions">Contributions</span>

**Bug reports, feature requests, and patches are always welcomed!**

That being said, take into account that:
* this is a very **personally tailored** tool, that fits just perfectly in my **personal workflow**; (note the "personal" used twice;) :)
* it tries to keep implementation complexity to a minimum;
  any feature that is implemented should either provide performance improvements (perceivable by a human),
  enable generally useful workflows (that otherwise are hard to implement via scripting),
  or increase flexibility (especially with regard to integration with one's generic tools);
* also see the anti-features listed in the [dedicated section](#anti-features);

Therefor, a feature request or patch might not be applied.
That doesn't mean the proposed idea is bad or worthless;  it just doesn't fit well with this tool.

However, given it is an open-source project, one can always just fork the project and take it in any direction.




--------------------------------------------------------------------------------




## <span id="license">Notice (copyright and licensing)</span>


### Notice -- short version

The code is licensed under GPL 3 or later.


### Notice -- long version

For details about the copyright and licensing, please consult the [`notice.txt`](./documentation/licensing/notice.txt) file in the [`documentation/licensing`](./documentation/licensing) folder.

If someone requires the sources and/or documentation to be released
under a different license, please send an email to the authors,
stating the licensing requirements, accompanied by the reasons
and other details; then, depending on the situation, the authors might
release the sources and/or documentation under a different license.

