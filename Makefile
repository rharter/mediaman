SELFPKG := github.com/rharter/mediaman
VERSION := 0.0.1
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
PKGS := \
database \
model
PKGS := $(addprefix github.com/rharter/mediaman/pkg/,$(PKGS))
.PHONY := test $(PKGS)

all: embed build

build:
	go build -o bin/mediaman -ldflags "-X main.version $(VERSION)dev-$(SHA)" $(SELFPKG)/cmd/mediaman

build-dist:
	godep go build -o bin/mediaman -ldflags "-X main.version $(VERSION)-$(SHA)" $(SELFPKG)/vmd/mediaman

bump-deps: deps

deps:
	go get -u -t -v ./...

# Embed static assets
embed: js rice
	cd cmd/mediaman && rice embed
	cd pkg/template && rice embed

js:
	#cd cmd/mediaman/assets && find js -name "*.js" ! -name '.*' ! -name "main.js" -exec cat {} \; > js/main.js

test: $(PKGS)

$(PKGS): godep
	godep go text -v $@

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
	go install github.com/GeertJohan/go.rice/rice

sample:
	mkdir sample
	touch "sample/12 Years a Slave [2013].idx"
	touch "sample/12 Years a Slave [2013].mkv"
	touch "sample/12 Years a Slave [2013].orig.nfo"
	touch "sample/12 Years a Slave [2013].sub"
	touch "sample/21 Jump Street [2012].en.sub"
	touch "sample/21 Jump Street [2012].idx"
	touch "sample/21 Jump Street [2012].mkv"
	touch "sample/21 Jump Street [2012].orig.nfo"
	touch "sample/22 Jump Street [2014].orig.nfo"
	touch "sample/22 Jump Street [2014].orig.txt"
	touch "sample/22 Jump Street [2014].wmv"
	touch "sample/After Earth (2013).mkv"
	touch "sample/After Earth (2013).mp4"
	touch "sample/After Earth (2013).orig.nfo"
	touch "sample/Alien [1979].mkv"
	touch "sample/Alien [1979].orig.nfo"
	touch "sample/American Hustle [2013].idx"
	touch "sample/American Hustle [2013].mkv"
	touch "sample/American Hustle [2013].orig.nfo"
	touch "sample/American Hustle [2013].sub"
	touch "sample/A Million Ways to Die in the West [2014].orig.nfo"
	touch "sample/A Million Ways to Die in the West [2014].orig.txt"
	touch "sample/A Million Ways to Die in the West [2014].wmv"
	touch "sample/Anchorman 2 The Legend Continues [2013].mkv"
	touch "sample/Bad Words [2014].avi"
	touch "sample/Bad Words [2014].orig.nfo"
	touch "sample/Bad Words [2014].orig.txt"
	touch "sample/Captain America The Winter Soldier [2014].orig.txt"
	touch "sample/Captain America The Winter Soldier [2014].wmv"
	touch "sample/Dawn of the Planet of the Apes [2014].orig.nfo"
	touch "sample/Dawn of the Planet of the Apes [2014].orig.txt"
	touch "sample/Dawn of the Planet of the Apes [2014].wmv"
	touch "sample/Delivery Man [2013].avi"
	touch "sample/Delivery Man [2013].idx"
	touch "sample/Delivery Man [2013].mkv"
	touch "sample/Delivery Man [2013].orig.nfo"
	touch "sample/Delivery Man [2013].orig.txt"
	touch "sample/Delivery Man [2013].sub"
	touch "sample/Despicable Me (2010).mkv"
	touch "sample/Despicable Me (2010).orig.nfo"
	touch "sample/Despicable Me 2 (2013).mkv"
	touch "sample/Despicable Me 2 (2013).orig.nfo"
	touch "sample/Divergent [2014].orig.nfo"
	touch "sample/Divergent [2014].orig.txt"
	touch "sample/Divergent [2014].wmv"
	touch "sample/Ender's Game [2013].mkv"
	touch "sample/Frozen [2013].mkv"
	touch "sample/Frozen [2013].orig.nfo"
	touch "sample/Guardians of the Galaxy [2014].orig.nfo"
	touch "sample/Guardians of the Galaxy [2014].orig.txt"
	touch "sample/Guardians of the Galaxy [2014].wmv"
	touch "sample/Identity Thief (2013).mkv"
	touch "sample/Identity Thief (2013).orig.nfo"
	touch "sample/In Time (2011).mkv"
	touch "sample/In Time (2011).orig.nfo"
	touch "sample/Magic Mike (2012).mkv"
	touch "sample/Monsters University (2013).mkv"
	touch "sample/Monsters University (2013).orig.nfo"
	touch "sample/Neighbors [2014].orig.nfo"
	touch "sample/Neighbors [2014].orig.txt"
	touch "sample/Neighbors [2014].wmv"
	touch "sample/Oblivion (2013).mkv"
	touch "sample/Oblivion (2013).orig.nfo"
	touch "sample/Predator [1987].mkv"
	touch "sample/Predator [1987].orig.nfo"
	touch "sample/Prometheus [2012].mkv"
	touch "sample/Prometheus [2012].orig.nfo"
	touch "sample/Rio 2 [2014].orig.nfo"
	touch "sample/Rio 2 [2014].orig.txt"
	touch "sample/Rio 2 [2014].wmv"
	touch "sample/Saving Mr. Banks [2013].avi"
	touch "sample/Saving Mr. Banks [2013].orig.nfo"
	touch "sample/Saving Mr. Banks [2013].orig.txt"
	touch "sample/Star Trek Into Darkness (2013).mkv"
	touch "sample/Star Trek Into Darkness (2013).orig.nfo"
	touch "sample/Star Wars Episode II - Attack of the Clones [2002].mkv"
	touch "sample/Star Wars Episode II- Attack of the Clones [2002].orig.nfo"
	touch "sample/Star Wars Episode III - Revenge of the Sith [2005].mkv"
	touch "sample/Star Wars Episode I - The Phantom Menace [1999].mkv"
	touch "sample/Star Wars Episode IV - A New Hope [1977].mkv"
	touch "sample/Star Wars Episode VI - Return of the Jedi [1983].avi"
	touch "sample/Star Wars Episode VI- Return of the Jedi [1983].orig.nfo"
	touch "sample/Star Wars Episode VI- Return of the Jedi [1983].orig.txt"
	touch "sample/Star Wars Episode V - The Empire Strikes Back [1980].mkv"
	touch "sample/Star Wars Episode V- The Empire Strikes Back [1980].orig.nfo"
	touch "sample/The Amazing Spider-Man [2012].mkv"
	touch "sample/The Amazing Spider-Man [2012].orig.nfo"
	touch "sample/The Amazing Spider-Man 2 [2014].orig.nfo"
	touch "sample/The Amazing Spider-Man 2 [2014].orig.txt"
	touch "sample/The Amazing Spider-Man 2 [2014].wmv"
	touch "sample/The Expendables 3 [2014].orig.nfo"
	touch "sample/The Expendables 3 [2014].orig.txt"
	touch "sample/The Expendables 3 [2014].wmv"
	touch "sample/The Golden Compass [2007].avi"
	touch "sample/The Golden Compass [2007].orig.nfo"
	touch "sample/The Golden Compass [2007].orig.txt"
	touch "sample/The Grand Budapest Hotel [2014].orig.txt"
	touch "sample/The Grand Budapest Hotel [2014].wmv"
	touch "sample/The Hobbit The Battle of the Five Armies [2014].avi"
	touch "sample/The Hobbit The Battle of the Five Armies [2014].orig.nfo"
	touch "sample/The Hobbit The Battle of the Five Armies [2014].orig.txt"
	touch "sample/The Hunger Games Catching Fire [2013].avi"
	touch "sample/The Hunger Games Mockingjay- Part 1 [2014].orig.nfo"
	touch "sample/The Hunger Games Mockingjay- Part 1 [2014].orig.txt"
	touch "sample/The Hunger Games Mockingjay - Part 1 [2014].wmv"
	touch "sample/The Legend of Hercules [2014].mkv"
	touch "sample/The LEGO Movie [2014].avi"
	touch "sample/The LEGO Movie [2014].orig.nfo"
	touch "sample/The LEGO Movie [2014].orig.txt"
	touch "sample/The Lunchbox [2013].mkv"
	touch "sample/The Matrix (1999).mkv"
	touch "sample/The Matrix (1999).orig.nfo"
	touch "sample/The Matrix (1999).srt"
	touch "sample/The Matrix Reloaded (2003).mkv"
	touch "sample/The Matrix Reloaded (2003).orig.nfo"
	touch "sample/The Matrix Reloaded (2003).srt"
	touch "sample/The Monuments Men [2014].avi"
	touch "sample/The Monuments Men [2014].orig.nfo"
	touch "sample/The Monuments Men [2014].orig.txt"
	touch "sample/Thor The Dark World (2013).mkv"
	touch "sample/Thor The Dark World (2013).orig.nfo"
	touch "sample/Transformers Age of Extinction [2014].orig.nfo"
	touch "sample/Transformers Age of Extinction [2014].orig.txt"
	touch "sample/Transformers Age of Extinction [2014].wmv"
	touch "sample/Up (2009).mkv"
	touch "sample/Up (2009).orig.nfo"
	touch "sample/Veronica Mars [2014].avi"
	touch "sample/Veronica Mars [2014].orig.nfo"
	touch "sample/Veronica Mars [2014].orig.txt"
	touch "sample/Warm Bodies [2013].en.sub"
	touch "sample/Warm Bodies [2013].es.sub"
	touch "sample/Warm Bodies [2013].idx"
	touch "sample/Warm Bodies [2013].mkv"
	touch "sample/Warm Bodies [2013].orig.nfo"
	touch "sample/World War Z [2013].idx"
	touch "sample/World War Z [2013].mkv"
	touch "sample/World War Z [2013].orig.nfo"
	touch "sample/World War Z [2013].sub"
	touch "sample/X-Men Days of Future Past [2014].orig.nfo"
	touch "sample/X-Men Days of Future Past [2014].orig.txt"
	touch "sample/X-Men Days of Future Past [2014].wmv"