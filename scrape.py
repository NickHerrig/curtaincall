import requests
from bs4 import BeautifulSoup
import re
from pprint import pprint

def main():
    '''Description of the function'''
    URL = "https://www.ibdb.com/Grosses/ViewProduction"
    PARAMS = {
    "IsNationalTour": True,
    "IsBroadway": False,
    "touringstartmonth": 10,
    "touringstartday": 16,
    "touringstartyear": 2021,
    "touringendmonth": 10,
    "touringendday": 17,
    "touringendyear": 2021,
    }

    res = requests.post(URL, params=PARAMS)
    soup = BeautifulSoup(res.content, "html.parser")
    links = soup.find_all("a", href=re.compile("tour-production"))
    pprint(links)

if __name__ == '__main__':
    main()

