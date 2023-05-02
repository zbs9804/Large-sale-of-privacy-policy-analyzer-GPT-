from bs4 import BeautifulSoup;
import requests
from google_play_scraper import app
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.common.proxy import *

# Top100 = webdriver.Chrome()
# Top100.get('https://www.appbrain.com/stats/google-play-rankings/top_free/health_fitness/us')
# simple_soup = BeautifulSoup(Top100.page_source, 'html.parser')

# locator = "body div.data-table-container table#rankings-table tbody tr td.ranking-app-cell"
# app_id = simple_soup.select(locator)
# app_ids = []
# for e in app_id:
#     locator = "a"
#     res = e.select_one(locator)
#     app_ids.append(res.attrs['href'].split('/')[3])

# def getLink(id):
#     result = app(
#         id,
#         lang='en', # defaults to 'en'
#         country='us' # defaults to 'us'
#     )
#     return(result['privacyPolicy'])

# counter = 1
# file = open('/Users/leonwang/Desktop/Privacy_policies.txt','w')
# for x in app_ids :
#     file.write(getLink(x)+"\n")
#     print(str(counter) + "Done")
#     counter += 1
# file.close()



def getAppUrl(html):
    res = html.select_one("a").attrs['href']
    urls.append(res)

def getPrivacyUrl(url):
    appleTop100.get(url)
    simple_soup = BeautifulSoup(appleTop100.page_source, 'html.parser')
    locator = "body div.ember-view main div.animation-wrapper.is-visible section.l-content-width.section.section--bordered.app-privacy p:nth-child(2) a"
    res = simple_soup.select_one(locator).attrs['href']
    return res

appleTop100 = webdriver.Chrome()
appleTop100.get('https://apps.apple.com/us/charts/iphone/health-fitness-apps/6013?chart=top-free')
simple_soup = BeautifulSoup(appleTop100.page_source, 'html.parser')

locator = "#charts-content-section ol li"
appHtml = simple_soup.select(locator)
urls = []
privacyUrls = []
# print(appHtml)
for x in appHtml:
    getAppUrl(x)
count = 0
applePrivacy = open('/Users/leonwang/Desktop/applePrivacy_policies.txt','w')
for url in urls:
    applePrivacy.write(getPrivacyUrl(url)+"\n")
    print(str(count) + "Done!")
    count += 1
