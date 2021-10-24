import requests
from bs4 import BeautifulSoup
import re
from pprint import pprint
import datetime
from collections import namedtuple
import re


SHOW_FIELDS = [
        'name',
        'songs',
        ]

Show = namedtuple("Show", SHOW_FIELDS)

SONG_FIELDS = [
        'name',
        'performer',
        'act',
        ]
Song = namedtuple("Song", SONG_FIELDS)

def get_songs(s):
    """Return a list of songs"""

    songs = []
    song_id = s.find("div", id=re.compile("Songs"))
    if song_id:
        song_divs = song_id.find_all("div", {"class": "col s6"})
        titles = [ div.get_text() for div in song_divs ]
        titles = [ " ".join(t.split())  for t in titles ]

        try:
            act_two_index = titles.index("Act 2")
        except ValueError:
            act_two_index = None

        if act_two_index:
            act_one_songs = titles[2:act_two_index]
            act_two_songs = titles[act_two_index+2:]
            for val in zip(act_one_songs[::2], act_one_songs[1::2]):
                s = Song(
                    val[0],
                    val[1],
                    "One",
                )
                songs.append(s)

            for val in zip(act_one_songs[::2], act_one_songs[1::2]):
                s = Song(
                    val[0],
                    val[1],
                    "Two",
                )
                songs.append(s)

            return songs

        for val in zip(titles[::2], titles[1::2]):
            s = Song(
                val[0],
                val[1],
                None
            )
            songs.append(s)

        return songs


def get_cast(s):
    print("parsing cast")
    #cast_id = s.find("div", id=re.compile("OpeningNightCast"))
    #cast = cast_id.find_all("div", {"class": "row mobile-a-align"})
    #for c in cast:
        #cast_mems = c.find_all("div", {"class": "col m4 s12"})
        #print(cast_mems)
        #cast =  cast_mems[0].get_text().strip()
        #role =  cast_mems[1].get_text().strip()

def main():

    ### Get Current Touring Shows ###
    now = datetime.datetime.now()
    last_week = now - datetime.timedelta(days=1)
    next_week = now + datetime.timedelta(days=1)
    URL = "https://www.ibdb.com/Grosses/ViewProduction"
    PARAMS = {
    "IsNationalTour": True,
    "IsBroadway": False,
    "touringstartmonth": last_week.month,
    "touringstartday": last_week.day,
    "touringstartyear": last_week.year,
    "touringendmonth": next_week.month,
    "touringendday": next_week.day,
    "touringendyear": next_week.year,
    }
    res = requests.post(URL, params=PARAMS)
    soup = BeautifulSoup(res.content, "html.parser")
    links = soup.find_all("a", href=re.compile("tour-production"))
    BASE = "https://www.ibdb.com"
    show_names = [ link.get_text() for link in links]
    show_links = [ BASE + link['href'] for link in links ]
    show_urls = dict(zip(show_names, show_links))

    shows = []
    for show, url in show_urls.items():
        songs_url = "".join([url, "#Songs"])
        res = requests.get(songs_url)
        soup = BeautifulSoup(res.content, "html.parser")
        songs = get_songs(soup)

        show = Show(
                show,
                songs,
                )

        shows.append(show)

    for show in shows:
        pprint(show.name)
        pprint(show.songs)


if __name__ == '__main__':
    main()

