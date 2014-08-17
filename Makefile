SELFPKG := github.com/rharter/mediaman
VERSION := 0.0.1
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
PKGS := \
database \
model \
api \
handler \
processor \
transcoding
PKGS := $(addprefix github.com/rharter/mediaman/pkg/,$(PKGS))
.PHONY := test $(PKGS)

all: embed build

build:
	go build -o bin/mediaman -ldflags "-X main.version $(VERSION)dev-$(SHA)" $(SELFPKG)/cmd/mediaman

build-dist:
	godep go build -o bin/mediaman -ldflags "-X main.version $(VERSION)-$(SHA)" $(SELFPKG)/vmd/mediaman

bump-deps: deps vendor

deps: 
	go get -u -t -v ./...

vendor: godep
	godep save ./...

# Embed static assets
embed: js rice
	cd cmd/mediaman && rice embed
	cd pkg/template && rice embed

js:
	#cd cmd/mediaman/assets && find js -name "*.js" ! -name '.*' ! -name "main.js" -exec cat {} \; > js/main.js

test:  $(PKGS)

$(PKGS):  godep
	godep go test -v $@

clean: rice
	cd cmd/mediaman && rice clean
	cd pkg/template && rice clean
	rm -rf cmd/mediaman/mediaman
	rm -rf cmd/mediaman/mediaman.sqlite
	rm -rf bin/mediaman
	rm -rf mediaman.sqlite
	rm -rf /tmp/mediaman.sqlite
	rm -rf sample

run:
	bin/mediaman --port=":8080" --datasource="mediaman.sqlite"

godep:
	go get github.com/tools/godep

rice:
	go get github.com/GeertJohan/go.rice/rice
	go build github.com/GeertJohan/go.rice/rice

sample:
	mkdir -p sample/movies
	touch "sample/movies/12 Years a Slave [2013].idx"
	touch "sample/movies/12 Years a Slave [2013].mkv"
	touch "sample/movies/12 Years a Slave [2013].orig.nfo"
	touch "sample/movies/12 Years a Slave [2013].sub"
	touch "sample/movies/21 Jump Street [2012].en.sub"
	touch "sample/movies/21 Jump Street [2012].idx"
	touch "sample/movies/21 Jump Street [2012].mkv"
	touch "sample/movies/21 Jump Street [2012].orig.nfo"
	touch "sample/movies/22 Jump Street [2014].orig.nfo"
	touch "sample/movies/22 Jump Street [2014].orig.txt"
	touch "sample/movies/22 Jump Street [2014].wmv"
	touch "sample/movies/After Earth (2013).mkv"
	touch "sample/movies/After Earth (2013).mp4"
	touch "sample/movies/After Earth (2013).orig.nfo"
	touch "sample/movies/Alien [1979].mkv"
	touch "sample/movies/Alien [1979].orig.nfo"
	touch "sample/movies/American Hustle [2013].idx"
	touch "sample/movies/American Hustle [2013].mkv"
	touch "sample/movies/American Hustle [2013].orig.nfo"
	touch "sample/movies/American Hustle [2013].sub"
	touch "sample/movies/A Million Ways to Die in the West [2014].orig.nfo"
	touch "sample/movies/A Million Ways to Die in the West [2014].orig.txt"
	touch "sample/movies/A Million Ways to Die in the West [2014].wmv"
	touch "sample/movies/Anchorman 2 The Legend Continues [2013].mkv"
	touch "sample/movies/Bad Words [2014].avi"
	touch "sample/movies/Bad Words [2014].orig.nfo"
	touch "sample/movies/Bad Words [2014].orig.txt"
	touch "sample/movies/Captain America The Winter Soldier [2014].orig.txt"
	touch "sample/movies/Captain America The Winter Soldier [2014].wmv"
	touch "sample/movies/Dawn of the Planet of the Apes [2014].orig.nfo"
	touch "sample/movies/Dawn of the Planet of the Apes [2014].orig.txt"
	touch "sample/movies/Dawn of the Planet of the Apes [2014].wmv"
	touch "sample/movies/Delivery Man [2013].avi"
	touch "sample/movies/Delivery Man [2013].idx"
	touch "sample/movies/Delivery Man [2013].mkv"
	touch "sample/movies/Delivery Man [2013].orig.nfo"
	touch "sample/movies/Delivery Man [2013].orig.txt"
	touch "sample/movies/Delivery Man [2013].sub"
	touch "sample/movies/Despicable Me (2010).mkv"
	touch "sample/movies/Despicable Me (2010).orig.nfo"
	touch "sample/movies/Despicable Me 2 (2013).mkv"
	touch "sample/movies/Despicable Me 2 (2013).orig.nfo"
	touch "sample/movies/Divergent [2014].orig.nfo"
	touch "sample/movies/Divergent [2014].orig.txt"
	touch "sample/movies/Divergent [2014].wmv"
	touch "sample/movies/Ender's Game [2013].mkv"
	touch "sample/movies/Frozen [2013].mkv"
	touch "sample/movies/Frozen [2013].orig.nfo"
	touch "sample/movies/Guardians of the Galaxy [2014].orig.nfo"
	touch "sample/movies/Guardians of the Galaxy [2014].orig.txt"
	touch "sample/movies/Guardians of the Galaxy [2014].wmv"
	touch "sample/movies/Identity Thief (2013).mkv"
	touch "sample/movies/Identity Thief (2013).orig.nfo"
	touch "sample/movies/In Time (2011).mkv"
	touch "sample/movies/In Time (2011).orig.nfo"
	touch "sample/movies/Magic Mike (2012).mkv"
	touch "sample/movies/Monsters University (2013).mkv"
	touch "sample/movies/Monsters University (2013).orig.nfo"
	touch "sample/movies/Neighbors [2014].orig.nfo"
	touch "sample/movies/Neighbors [2014].orig.txt"
	touch "sample/movies/Neighbors [2014].wmv"
	touch "sample/movies/Oblivion (2013).mkv"
	touch "sample/movies/Oblivion (2013).orig.nfo"
	touch "sample/movies/Predator [1987].mkv"
	touch "sample/movies/Predator [1987].orig.nfo"
	touch "sample/movies/Prometheus [2012].mkv"
	touch "sample/movies/Prometheus [2012].orig.nfo"
	touch "sample/movies/Rio 2 [2014].orig.nfo"
	touch "sample/movies/Rio 2 [2014].orig.txt"
	touch "sample/movies/Rio 2 [2014].wmv"
	touch "sample/movies/Saving Mr. Banks [2013].avi"
	touch "sample/movies/Saving Mr. Banks [2013].orig.nfo"
	touch "sample/movies/Saving Mr. Banks [2013].orig.txt"
	touch "sample/movies/Star Trek Into Darkness (2013).mkv"
	touch "sample/movies/Star Trek Into Darkness (2013).orig.nfo"
	touch "sample/movies/Star Wars Episode II - Attack of the Clones [2002].mkv"
	touch "sample/movies/Star Wars Episode II- Attack of the Clones [2002].orig.nfo"
	touch "sample/movies/Star Wars Episode III - Revenge of the Sith [2005].mkv"
	touch "sample/movies/Star Wars Episode I - The Phantom Menace [1999].mkv"
	touch "sample/movies/Star Wars Episode IV - A New Hope [1977].mkv"
	touch "sample/movies/Star Wars Episode VI - Return of the Jedi [1983].avi"
	touch "sample/movies/Star Wars Episode VI- Return of the Jedi [1983].orig.nfo"
	touch "sample/movies/Star Wars Episode VI- Return of the Jedi [1983].orig.txt"
	touch "sample/movies/Star Wars Episode V - The Empire Strikes Back [1980].mkv"
	touch "sample/movies/Star Wars Episode V- The Empire Strikes Back [1980].orig.nfo"
	touch "sample/movies/The Amazing Spider-Man [2012].mkv"
	touch "sample/movies/The Amazing Spider-Man [2012].orig.nfo"
	touch "sample/movies/The Amazing Spider-Man 2 [2014].orig.nfo"
	touch "sample/movies/The Amazing Spider-Man 2 [2014].orig.txt"
	touch "sample/movies/The Amazing Spider-Man 2 [2014].wmv"
	touch "sample/movies/The Expendables 3 [2014].orig.nfo"
	touch "sample/movies/The Expendables 3 [2014].orig.txt"
	touch "sample/movies/The Expendables 3 [2014].wmv"
	touch "sample/movies/The Golden Compass [2007].avi"
	touch "sample/movies/The Golden Compass [2007].orig.nfo"
	touch "sample/movies/The Golden Compass [2007].orig.txt"
	touch "sample/movies/The Grand Budapest Hotel [2014].orig.txt"
	touch "sample/movies/The Grand Budapest Hotel [2014].wmv"
	touch "sample/movies/The Hobbit The Battle of the Five Armies [2014].avi"
	touch "sample/movies/The Hobbit The Battle of the Five Armies [2014].orig.nfo"
	touch "sample/movies/The Hobbit The Battle of the Five Armies [2014].orig.txt"
	touch "sample/movies/The Hunger Games Catching Fire [2013].avi"
	touch "sample/movies/The Hunger Games Mockingjay- Part 1 [2014].orig.nfo"
	touch "sample/movies/The Hunger Games Mockingjay- Part 1 [2014].orig.txt"
	touch "sample/movies/The Hunger Games Mockingjay - Part 1 [2014].wmv"
	touch "sample/movies/The Legend of Hercules [2014].mkv"
	touch "sample/movies/The LEGO Movie [2014].avi"
	touch "sample/movies/The LEGO Movie [2014].orig.nfo"
	touch "sample/movies/The LEGO Movie [2014].orig.txt"
	touch "sample/movies/The Lunchbox [2013].mkv"
	touch "sample/movies/The Matrix (1999).mkv"
	touch "sample/movies/The Matrix (1999).orig.nfo"
	touch "sample/movies/The Matrix (1999).srt"
	touch "sample/movies/The Matrix Reloaded (2003).mkv"
	touch "sample/movies/The Matrix Reloaded (2003).orig.nfo"
	touch "sample/movies/The Matrix Reloaded (2003).srt"
	touch "sample/movies/The Monuments Men [2014].avi"
	touch "sample/movies/The Monuments Men [2014].orig.nfo"
	touch "sample/movies/The Monuments Men [2014].orig.txt"
	touch "sample/movies/Thor The Dark World (2013).mkv"
	touch "sample/movies/Thor The Dark World (2013).orig.nfo"
	touch "sample/movies/Transformers Age of Extinction [2014].orig.nfo"
	touch "sample/movies/Transformers Age of Extinction [2014].orig.txt"
	touch "sample/movies/Transformers Age of Extinction [2014].wmv"
	touch "sample/movies/Up (2009).mkv"
	touch "sample/movies/Up (2009).orig.nfo"
	touch "sample/movies/Veronica Mars [2014].avi"
	touch "sample/movies/Veronica Mars [2014].orig.nfo"
	touch "sample/movies/Veronica Mars [2014].orig.txt"
	touch "sample/movies/Warm Bodies [2013].en.sub"
	touch "sample/movies/Warm Bodies [2013].es.sub"
	touch "sample/movies/Warm Bodies [2013].idx"
	touch "sample/movies/Warm Bodies [2013].mkv"
	touch "sample/movies/Warm Bodies [2013].orig.nfo"
	touch "sample/movies/World War Z [2013].idx"
	touch "sample/movies/World War Z [2013].mkv"
	touch "sample/movies/World War Z [2013].orig.nfo"
	touch "sample/movies/World War Z [2013].sub"
	touch "sample/movies/X-Men Days of Future Past [2014].orig.nfo"
	touch "sample/movies/X-Men Days of Future Past [2014].orig.txt"
	touch "sample/movies/X-Men Days of Future Past [2014].wmv"