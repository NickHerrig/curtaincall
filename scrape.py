import requests
from bs4 import BeautifulSoup
import re
from pprint import pprint
import datetime

def main():
    '''Description of the function'''
#    now = datetime.datetime.now()
#    last_week = now - datetime.timedelta(days=7)
#    next_week = now + datetime.timedelta(days=7)
#
#    URL = "https://www.ibdb.com/Grosses/ViewProduction"
#    PARAMS = {
#    "IsNationalTour": True,
#    "IsBroadway": False,
#    "touringstartmonth": last_week.month,
#    "touringstartday": last_week.day,
#    "touringstartyear": last_week.year,
#    "touringendmonth": next_week.month,
#    "touringendday": next_week.day,
#    "touringendyear": next_week.year,
#    }
#
#    res = requests.post(URL, params=PARAMS)
#    soup = BeautifulSoup(res.content, "html.parser")
#    links = soup.find_all("a", href=re.compile("tour-production"))
#
#    BASE = "https://www.ibdb.com"
#    show_links = [ BASE + link['href'] for link in links ]
#    pprint(show_links)

    SHOW_URL = "https://www.ibdb.com/tour-production/come-from-away-518408#Songs"
    res = requests.get(SHOW_URL)
    soup = BeautifulSoup(res.content, "html.parser")
    song_id = soup.find("div", id=re.compile("Songs"))
    songs = song_id.find_all("div", {"class": "col s6"})
    title = [ song.text.strip() for song in songs ]
    song_performers = tuple(zip(title[::2], title[1::2]))
    pprint(song_performers)

if __name__ == '__main__':
    main()

