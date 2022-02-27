package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

func TestParseItem(t *testing.T) {
	s := NewScraper([]string{}, 0.0087)

	html := `<div class="Products Products--grid Products--immersive js-immersive" data-logintype="notLogin" data-yid="" data-done="https%3A%2F%2Fauctions.yahoo.co.jp%2Fsearch%2Fsearch%3Fp%3D%25E3%2583%2594%25E3%2582%25AB%25E3%2583%2581%25E3%2583%25A5%25E3%2582%25A6%26va%3D%25E3%2583%2594%25E3%2582%25AB%25E3%2583%2581%25E3%2583%25A5%25E3%2582%25A6%26exflg%3D1%26b%3D1%26n%3D50%26s1%3Dnew%26o1%3Dd%26mode%3D4" data-prefecturecode="13" data-prefecturename="東京都" data-crumb="cf65aa852d2b328e5a11ce383d719eee8bc3f9ccce5bef13b9ce6034ebf513f7" data-bucketid="ctrl100" data-immersive-id="0">
    <div class="Products__list js-immersive-list">
                <ul class="Products__items">
                                                <li class="Product js-immersive-item">
                <div class="Product__image">
                    <a class="Product__imageLink js-immersive-itemLink js-rapid-override rapidnofollow js-browseHistory-add" data-auction-id="s1039793076" data-auction-category="2084259479" data-auction-title="【ポケモンカード】200/SM-P PROMO ピカチュウ プロモ　中古　ポケカ" data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/11618ebeaf2d01388cd527d397fc5b25df865b06/i-img480x640-1645952174wmwm5x342545.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="500" data-auction-isfreeshipping="" href="https://page.auctions.yahoo.co.jp/jp/auction/s1039793076" data-itemurl="https://auctions.yahoo.co.jp/show/auctionjs?aid=s1039793076&amp;_crumb=cf65aa852d2b328e5a11ce383d719eee8bc3f9ccce5bef13b9ce6034ebf513f7" data-ylk="rsec:aal;slk:ic;pos:1;atax:0;arbn:;catid:2084259479;cid:s1039793076;st:1645952230;end:1646312004;prat:0;op:;grat:99.6;ppstr:;cprate:;seltyp:0;best:" data-immersive-item-id="0" data-recommendurl="https://auctions.yahoo.co.jp/show/recommendjs?rid=41&amp;aid=s1039793076&amp;title=%E3%80%90%E3%83%9D%E3%82%B1%E3%83%A2%E3%83%B3%E3%82%AB%E3%83%BC%E3%83%89%E3%80%91200%2FSM-P+PROMO+%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6+%E3%83%97%E3%83%AD%E3%83%A2%E3%80%80%E4%B8%AD%E5%8F%A4%E3%80%80%E3%83%9D%E3%82%B1%E3%82%AB&amp;catid=2084259479&amp;cpath=0,25464,27727,25826,2084241343,2084259479&amp;is_adult_category=&amp;_crumb=dc26c18ae1af65dbdec3f7487d0fbb678840a5c55849aa23b4ae636161cf8593" data-rapid_p="398">
                                                                        <img class="Product__imageData" src="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/11618ebeaf2d01388cd527d397fc5b25df865b06/i-img480x640-1645952174wmwm5x342545.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" alt="【ポケモンカード】200/SM-P PROMO ピカチュウ プロモ　中古　ポケカ" width="75" height="100" loading="lazy">
                                                                                    </a>
                    <div class="HideSeller">
                                            <a class="HideSeller__link js-rapid-override" data-auction-id="s1039793076" data-auction-category="2084259479" href="https://auctions.yahoo.co.jp/search/search?p=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;va=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;exflg=1&amp;exsid=anonymous1&amp;b=1&amp;n=50&amp;s1=new&amp;o1=d&amp;mode=4" rel="nofollow" data-ylk="rsec:aal;slk:dslc;pos:1;arbn:" data-rapid_p="399">非表示
                            <div class="HideSeller__balloon">
                                <div class="HideSeller__detail">
                                    <p class="HideSeller__message">この出品者の商品を非表示にする</p>
                                </div>
                            </div>
                        </a>
                                        </div><!-- ./HideSeller -->
                                            <span class="Product__time">4日</span>
                                                                                            <div class="Product__button js-watch-button notlogin">
                                <a class="Button Button--watch is-off rapid-noclick-resp cl-noclick-log js-watch-sync" rel="nofollow" href="#dummy" data-deleteurl="https://auctions.yahoo.co.jp/watchlist/delete?aID=s1039793076&amp;.crumb=" data-addurl="https://auctions.yahoo.co.jp/watchlist/add?aID=s1039793076&amp;.crumb=" data-loading="0" data-operation="add" data-aid="s1039793076" data-ylk="rsec:aal;slk:wc;pos:1;atax:0;catid:2084259479;cid:s1039793076;st:1645952230;end:1646312004;prat:0;spmwl:0;op:;arbn:;sw:off;grat:99.6;ppstr:;cprate:;seltyp:0;best:" data-rapid_p="400">ウォッチ</a>
                                                                </div>
                                                <div class="Layer js-watch-button-layer"></div>
                                    </div><!-- /.Product__image -->
                <div class="Product__detail">
                    <div class="Product__bonus " data-auction-id="s1039793076" data-auction-endtime="1646312004" data-auction-buynowprice="0" data-auction-categoryidpath="0,25464,27727,25826,2084241343,2084259479" data-auction-caneasypayment="1" data-auction-price="500" data-auction-sellerid="anonymous1" data-auction-startprice="500" data-auction-isshoppingitem=""></div>
                    <h3 class="Product__title">
                        <a class="Product__titleLink js-rapid-override js-browseHistory-add" data-auction-id="s1039793076" data-auction-category="2084259479" data-auction-title="【ポケモンカード】200/SM-P PROMO ピカチュウ プロモ　中古　ポケカ" data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/11618ebeaf2d01388cd527d397fc5b25df865b06/i-img480x640-1645952174wmwm5x342545.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="500" data-auction-isfreeshipping="" href="https://page.auctions.yahoo.co.jp/jp/auction/s1039793076" target="_blank" rel="noopener" data-ylk="rsec:aal;slk:tc;pos:1;atax:0;arbn:;catid:2084259479;cid:s1039793076;etc:p=500,etm=1646312004,stm=1645952230;st:1645952230;end:1646312004;prat:0;op:;grat:99.6;ppstr:;cprate:;seltyp:0;best:" title="【ポケモンカード】200/SM-P PROMO ピカチュウ プロモ　中古　ポケカ" data-rapid_p="401">【ポケモンカード】200/SM-P PROMO ピカチュウ プロモ　中古　ポケカ</a>
                    </h3>
                     <span class="Product__price">
                        <span class="Product__priceValue u-textRed">500円</span>
                    </span><!-- ./Product__price -->
                                            <span class="Product__price">
                            <span class="Product__label">即決</span>
                                                            <span class="Product__priceValue">-</span>
                                                    </span><!-- ./Product__price -->
                                                        </div><!-- /.Product__detail -->
            </li>
                    <li class="Product js-immersive-item">
                <div class="Product__image">
                    <a class="Product__imageLink js-immersive-itemLink js-rapid-override rapidnofollow js-browseHistory-add" data-auction-id="s1039791085" data-auction-category="2084309054" data-auction-title="ポケモンカード　ボスの指令　(s9)(s4a)(s8b) フュージョンアーツ　VMAX スターバース　ピカチュウ　ミュウ　リザードン" data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/ddf5f4332f4de59492634c20a947430ebc77a5d0/i-img898x1198-1645951917ctrgi9219945.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="1000" data-auction-isfreeshipping="1" href="https://page.auctions.yahoo.co.jp/jp/auction/s1039791085" data-itemurl="https://auctions.yahoo.co.jp/show/auctionjs?aid=s1039791085&amp;_crumb=cf65aa852d2b328e5a11ce383d719eee8bc3f9ccce5bef13b9ce6034ebf513f7" data-ylk="rsec:aal;slk:ic;pos:2;atax:0;arbn:;catid:2084309054;cid:s1039791085;st:1645951918;end:1646563918;prat:0;op:;grat:100.0;ppstr:;cprate:;seltyp:0;best:" data-immersive-item-id="1" data-recommendurl="https://auctions.yahoo.co.jp/show/recommendjs?rid=41&amp;aid=s1039791085&amp;title=%E3%83%9D%E3%82%B1%E3%83%A2%E3%83%B3%E3%82%AB%E3%83%BC%E3%83%89%E3%80%80%E3%83%9C%E3%82%B9%E3%81%AE%E6%8C%87%E4%BB%A4%E3%80%80%28s9%29%28s4a%29%28s8b%29+%E3%83%95%E3%83%A5%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B3%E3%82%A2%E3%83%BC%E3%83%84%E3%80%80VMAX+%E3%82%B9%E3%82%BF%E3%83%BC%E3%83%90%E3%83%BC%E3%82%B9%E3%80%80%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6%E3%80%80%E3%83%9F%E3%83%A5%E3%82%A6%E3%80%80%E3%83%AA%E3%82%B6%E3%83%BC%E3%83%89%E3%83%B3&amp;catid=2084309054&amp;cpath=0,25464,27727,25826,2084241343,2084309054&amp;is_adult_category=&amp;_crumb=dc26c18ae1af65dbdec3f7487d0fbb678840a5c55849aa23b4ae636161cf8593" data-rapid_p="402">
                                                                        <img class="Product__imageData" src="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/ddf5f4332f4de59492634c20a947430ebc77a5d0/i-img898x1198-1645951917ctrgi9219945.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" alt="ポケモンカード　ボスの指令　(s9)(s4a)(s8b) フュージョンアーツ　VMAX スターバース　ピカチュウ　ミュウ　リザードン" width="74" height="100" loading="lazy">
                                                                                        <span class="Product__icon Product__icon--freeShipping">送料無料</span>
                                        </a>
                    <div class="HideSeller">
                                            <a class="HideSeller__link js-rapid-override" data-auction-id="s1039791085" data-auction-category="2084309054" href="https://auctions.yahoo.co.jp/search/search?p=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;va=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;exflg=1&amp;exsid=anonymous2&amp;b=1&amp;n=50&amp;s1=new&amp;o1=d&amp;mode=4" rel="nofollow" data-ylk="rsec:aal;slk:dslc;pos:2;arbn:" data-rapid_p="403">非表示
                            <div class="HideSeller__balloon">
                                <div class="HideSeller__detail">
                                    <p class="HideSeller__message">この出品者の商品を非表示にする</p>
                                </div>
                            </div>
                        </a>
                                        </div><!-- ./HideSeller -->
                                            <span class="Product__time">7日</span>
                                                                                            <div class="Product__button js-watch-button notlogin">
                                <a class="Button Button--watch is-off rapid-noclick-resp cl-noclick-log js-watch-sync" rel="nofollow" href="#dummy" data-deleteurl="https://auctions.yahoo.co.jp/watchlist/delete?aID=s1039791085&amp;.crumb=" data-addurl="https://auctions.yahoo.co.jp/watchlist/add?aID=s1039791085&amp;.crumb=" data-loading="0" data-operation="add" data-aid="s1039791085" data-ylk="rsec:aal;slk:wc;pos:2;atax:0;catid:2084309054;cid:s1039791085;st:1645951918;end:1646563918;prat:0;spmwl:0;op:;arbn:;sw:off;grat:100.0;ppstr:;cprate:;seltyp:0;best:" data-rapid_p="404">ウォッチ</a>
                                                                </div>
                                                <div class="Layer js-watch-button-layer"></div>
                                    </div><!-- /.Product__image -->
                <div class="Product__detail">
                    <div class="Product__bonus " data-auction-id="s1039791085" data-auction-endtime="1646563918" data-auction-buynowprice="0" data-auction-categoryidpath="0,25464,27727,25826,2084241343,2084309054" data-auction-caneasypayment="1" data-auction-price="1000" data-auction-sellerid="anonymous2" data-auction-startprice="1000" data-auction-isshoppingitem=""></div>
                    <h3 class="Product__title">
                        <a class="Product__titleLink js-rapid-override js-browseHistory-add" data-auction-id="s1039791085" data-auction-category="2084309054" data-auction-title="ポケモンカード　ボスの指令　(s9)(s4a)(s8b) フュージョンアーツ　VMAX スターバース　ピカチュウ　ミュウ　リザードン" data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/ddf5f4332f4de59492634c20a947430ebc77a5d0/i-img898x1198-1645951917ctrgi9219945.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="1000" data-auction-isfreeshipping="1" href="https://page.auctions.yahoo.co.jp/jp/auction/s1039791085" target="_blank" rel="noopener" data-ylk="rsec:aal;slk:tc;pos:2;atax:0;arbn:;catid:2084309054;cid:s1039791085;etc:p=1000,etm=1646563918,stm=1645951918;st:1645951918;end:1646563918;prat:0;op:;grat:100.0;ppstr:;cprate:;seltyp:0;best:" title="ポケモンカード　ボスの指令　(s9)(s4a)(s8b) フュージョンアーツ　VMAX スターバース　ピカチュウ　ミュウ　リザードン" data-rapid_p="405">ポケモンカード　ボスの指令　(s9)(s4a)(s8b) フュージョンアーツ　VMAX スターバース　ピカチュウ　ミュウ　リザードン</a>
                    </h3>
                     <span class="Product__price">
                        <span class="Product__priceValue u-textRed">1,000円</span>
                    </span><!-- ./Product__price -->
                                            <span class="Product__price">
                            <span class="Product__label">即決</span>
                                                            <span class="Product__priceValue">-</span>
                                                    </span><!-- ./Product__price -->
                                                        </div><!-- /.Product__detail -->
            </li>
                    <li class="Product js-immersive-item">
                <div class="Product__image">
                    <a class="Product__imageLink js-immersive-itemLink js-rapid-override rapidnofollow js-browseHistory-add" data-auction-id="r1039797461" data-auction-category="26082" data-auction-title="ぬいぐるみ Pokmon Cafe パティシエピカチュウ" data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/063d38915bd7c71f574adf8761c160b009998b9b/i-img483x483-1645951569fpgcyx218826.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="5278" data-auction-isfreeshipping="1" href="https://page.auctions.yahoo.co.jp/jp/auction/r1039797461" data-itemurl="https://auctions.yahoo.co.jp/show/auctionjs?aid=r1039797461&amp;_crumb=cf65aa852d2b328e5a11ce383d719eee8bc3f9ccce5bef13b9ce6034ebf513f7" data-ylk="rsec:aal;slk:ic;pos:3;atax:0;arbn:;catid:26082;cid:r1039797461;st:1645951576;end:1646059555;prat:0;op:;grat:;ppstr:;cprate:;seltyp:2;best:" data-immersive-item-id="2" data-recommendurl="https://auctions.yahoo.co.jp/show/recommendjs?rid=41&amp;aid=r1039797461&amp;title=%E3%81%AC%E3%81%84%E3%81%90%E3%82%8B%E3%81%BF+Pokmon+Cafe+%E3%83%91%E3%83%86%E3%82%A3%E3%82%B7%E3%82%A8%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;catid=26082&amp;cpath=0,25464,26082&amp;is_adult_category=&amp;_crumb=dc26c18ae1af65dbdec3f7487d0fbb678840a5c55849aa23b4ae636161cf8593" data-rapid_p="406">
                                                                        <img class="Product__imageData" src="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/063d38915bd7c71f574adf8761c160b009998b9b/i-img483x483-1645951569fpgcyx218826.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" alt="ぬいぐるみ Pokmon Cafe パティシエピカチュウ" width="100" height="100" loading="lazy">
                                                                                        <span class="Product__icon Product__icon--freeShipping">送料無料</span>
                                        </a>
                    <div class="HideSeller">
                                            <a class="HideSeller__link js-rapid-override" data-auction-id="r1039797461" data-auction-category="26082" href="https://auctions.yahoo.co.jp/search/search?p=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;va=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;exflg=1&amp;exsid=anonymous3&amp;b=1&amp;n=50&amp;s1=new&amp;o1=d&amp;mode=4" rel="nofollow" data-ylk="rsec:aal;slk:dslc;pos:3;arbn:" data-rapid_p="407">非表示
                            <div class="HideSeller__balloon">
                                <div class="HideSeller__detail">
                                    <p class="HideSeller__message">この出品者の商品を非表示にする</p>
                                </div>
                            </div>
                        </a>
                                        </div><!-- ./HideSeller -->
                                            <span class="Product__time">1日</span>
                                                                                            <div class="Product__button js-watch-button notlogin">
                                <a class="Button Button--watch is-off rapid-noclick-resp cl-noclick-log js-watch-sync" rel="nofollow" href="#dummy" data-deleteurl="https://auctions.yahoo.co.jp/watchlist/delete?aID=r1039797461&amp;.crumb=" data-addurl="https://auctions.yahoo.co.jp/watchlist/add?aID=r1039797461&amp;.crumb=" data-loading="0" data-operation="add" data-aid="r1039797461" data-ylk="rsec:aal;slk:wc;pos:3;atax:0;catid:26082;cid:r1039797461;st:1645951576;end:1646059555;prat:0;spmwl:0;op:;arbn:;sw:off;grat:;ppstr:;cprate:;seltyp:2;best:" data-rapid_p="408">ウォッチ</a>
                                                                </div>
                                                <div class="Layer js-watch-button-layer"></div>
                                    </div><!-- /.Product__image -->
                <div class="Product__detail">
                    <div class="Product__bonus " data-auction-id="r1039797461" data-auction-endtime="1646059555" data-auction-buynowprice="5278" data-auction-categoryidpath="0,25464,26082" data-auction-caneasypayment="1" data-auction-price="5278" data-auction-sellerid="anonymous3" data-auction-startprice="5278" data-auction-isshoppingitem=""></div>
                    <h3 class="Product__title">
                        <a class="Product__titleLink js-rapid-override js-browseHistory-add" data-auction-id="r1039797461" data-auction-category="26082" data-auction-title="ぬいぐるみ Pokmon Cafe パティシエピカチュウ" data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/063d38915bd7c71f574adf8761c160b009998b9b/i-img483x483-1645951569fpgcyx218826.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="5278" data-auction-isfreeshipping="1" href="https://page.auctions.yahoo.co.jp/jp/auction/r1039797461" target="_blank" rel="noopener" data-ylk="rsec:aal;slk:tc;pos:3;atax:0;arbn:;catid:26082;cid:r1039797461;etc:p=5278,etm=1646059555,stm=1645951576;st:1645951576;end:1646059555;prat:0;op:;grat:;ppstr:;cprate:;seltyp:2;best:" title="ぬいぐるみ Pokmon Cafe パティシエピカチュウ" data-rapid_p="409">ぬいぐるみ Pokmon Cafe パティシエピカチュウ</a>
                    </h3>
                     <span class="Product__price">
                        <span class="Product__priceValue u-textRed">5,278円</span>
                    </span><!-- ./Product__price -->
                                            <span class="Product__price">
                            <span class="Product__label">即決</span>
                                                            <span class="Product__priceValue">5,278円</span>
                                                    </span><!-- ./Product__price -->
                                                        </div><!-- /.Product__detail -->
            </li>
                    <li class="Product js-immersive-item">
                <div class="Product__image">
                    <a class="Product__imageLink js-immersive-itemLink js-rapid-override rapidnofollow js-browseHistory-add" data-auction-id="x1039785922" data-auction-category="2084259479" data-auction-title="【ポケモンカード】283/SM-P ヨコハマのピカチュウ プロモ　中古　ポケカ" data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/11618ebeaf2d01388cd527d397fc5b25df865b06/i-img480x640-1645951481sibp7a139407.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="1000" data-auction-isfreeshipping="" href="https://page.auctions.yahoo.co.jp/jp/auction/x1039785922" data-itemurl="https://auctions.yahoo.co.jp/show/auctionjs?aid=x1039785922&amp;_crumb=cf65aa852d2b328e5a11ce383d719eee8bc3f9ccce5bef13b9ce6034ebf513f7" data-ylk="rsec:aal;slk:ic;pos:4;atax:0;arbn:;catid:2084259479;cid:x1039785922;st:1645951531;end:1646401379;prat:0;op:;grat:99.6;ppstr:;cprate:;seltyp:0;best:" data-immersive-item-id="3" data-recommendurl="https://auctions.yahoo.co.jp/show/recommendjs?rid=41&amp;aid=x1039785922&amp;title=%E3%80%90%E3%83%9D%E3%82%B1%E3%83%A2%E3%83%B3%E3%82%AB%E3%83%BC%E3%83%89%E3%80%91283%2FSM-P+%E3%83%A8%E3%82%B3%E3%83%8F%E3%83%9E%E3%81%AE%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6+%E3%83%97%E3%83%AD%E3%83%A2%E3%80%80%E4%B8%AD%E5%8F%A4%E3%80%80%E3%83%9D%E3%82%B1%E3%82%AB&amp;catid=2084259479&amp;cpath=0,25464,27727,25826,2084241343,2084259479&amp;is_adult_category=&amp;_crumb=dc26c18ae1af65dbdec3f7487d0fbb678840a5c55849aa23b4ae636161cf8593" data-rapid_p="410">
                                                                        <img class="Product__imageData" src="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/11618ebeaf2d01388cd527d397fc5b25df865b06/i-img480x640-1645951481sibp7a139407.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" alt="【ポケモンカード】283/SM-P ヨコハマのピカチュウ プロモ　中古　ポケカ" width="75" height="100" loading="lazy">
                                                                                    </a>
                    <div class="HideSeller">
                                            <a class="HideSeller__link js-rapid-override" data-auction-id="x1039785922" data-auction-category="2084259479" href="https://auctions.yahoo.co.jp/search/search?p=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;va=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;exflg=1&amp;exsid=anonymous1&amp;b=1&amp;n=50&amp;s1=new&amp;o1=d&amp;mode=4" rel="nofollow" data-ylk="rsec:aal;slk:dslc;pos:4;arbn:" data-rapid_p="411">非表示
                            <div class="HideSeller__balloon">
                                <div class="HideSeller__detail">
                                    <p class="HideSeller__message">この出品者の商品を非表示にする</p>
                                </div>
                            </div>
                        </a>
                                        </div><!-- ./HideSeller -->
                                            <span class="Product__time">5日</span>
                                                                                            <div class="Product__button js-watch-button notlogin">
                                <a class="Button Button--watch is-off rapid-noclick-resp cl-noclick-log js-watch-sync" rel="nofollow" href="#dummy" data-deleteurl="https://auctions.yahoo.co.jp/watchlist/delete?aID=x1039785922&amp;.crumb=" data-addurl="https://auctions.yahoo.co.jp/watchlist/add?aID=x1039785922&amp;.crumb=" data-loading="0" data-operation="add" data-aid="x1039785922" data-ylk="rsec:aal;slk:wc;pos:4;atax:0;catid:2084259479;cid:x1039785922;st:1645951531;end:1646401379;prat:0;spmwl:0;op:;arbn:;sw:off;grat:99.6;ppstr:;cprate:;seltyp:0;best:" data-rapid_p="412">ウォッチ</a>
                                                                </div>
                                                <div class="Layer js-watch-button-layer"></div>
                                    </div><!-- /.Product__image -->
                <div class="Product__detail">
                    <div class="Product__bonus " data-auction-id="x1039785922" data-auction-endtime="1646401379" data-auction-buynowprice="0" data-auction-categoryidpath="0,25464,27727,25826,2084241343,2084259479" data-auction-caneasypayment="1" data-auction-price="1000" data-auction-sellerid="anonymous1" data-auction-startprice="1000" data-auction-isshoppingitem=""></div>
                    <h3 class="Product__title">
                        <a class="Product__titleLink js-rapid-override js-browseHistory-add" data-auction-id="x1039785922" data-auction-category="2084259479" data-auction-title="【ポケモンカード】283/SM-P ヨコハマのピカチュウ プロモ　中古　ポケカ" data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/11618ebeaf2d01388cd527d397fc5b25df865b06/i-img480x640-1645951481sibp7a139407.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="1000" data-auction-isfreeshipping="" href="https://page.auctions.yahoo.co.jp/jp/auction/x1039785922" target="_blank" rel="noopener" data-ylk="rsec:aal;slk:tc;pos:4;atax:0;arbn:;catid:2084259479;cid:x1039785922;etc:p=1000,etm=1646401379,stm=1645951531;st:1645951531;end:1646401379;prat:0;op:;grat:99.6;ppstr:;cprate:;seltyp:0;best:" title="【ポケモンカード】283/SM-P ヨコハマのピカチュウ プロモ　中古　ポケカ" data-rapid_p="413">【ポケモンカード】283/SM-P ヨコハマのピカチュウ プロモ　中古　ポケカ</a>
                    </h3>
                     <span class="Product__price">
                        <span class="Product__priceValue u-textRed">1,000円</span>
                    </span><!-- ./Product__price -->
                                            <span class="Product__price">
                            <span class="Product__label">即決</span>
                                                            <span class="Product__priceValue">-</span>
                                                    </span><!-- ./Product__price -->
                                                        </div><!-- /.Product__detail -->
            </li>
                    <li class="Product js-immersive-item">
                <div class="Product__image">
                    <a class="Product__imageLink js-immersive-itemLink js-rapid-override rapidnofollow js-browseHistory-add" data-auction-id="s1039795118" data-auction-category="2084259479" data-auction-title="☆ポケモンカード☆旧裏☆ピカチュウ " data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/1198328c2d101a68702a8996f2254cce6da0456c/i-img900x1200-1645951291o3mldo875826.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="500" data-auction-isfreeshipping="" href="https://page.auctions.yahoo.co.jp/jp/auction/s1039795118" data-itemurl="https://auctions.yahoo.co.jp/show/auctionjs?aid=s1039795118&amp;_crumb=cf65aa852d2b328e5a11ce383d719eee8bc3f9ccce5bef13b9ce6034ebf513f7" data-ylk="rsec:aal;slk:ic;pos:5;atax:0;arbn:;catid:2084259479;cid:s1039795118;st:1645951292;end:1646059292;prat:0;op:;grat:99.4;ppstr:;cprate:;seltyp:0;best:" data-immersive-item-id="4" data-recommendurl="https://auctions.yahoo.co.jp/show/recommendjs?rid=41&amp;aid=s1039795118&amp;title=%E2%98%86%E3%83%9D%E3%82%B1%E3%83%A2%E3%83%B3%E3%82%AB%E3%83%BC%E3%83%89%E2%98%86%E6%97%A7%E8%A3%8F%E2%98%86%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6+&amp;catid=2084259479&amp;cpath=0,25464,27727,25826,2084241343,2084259479&amp;is_adult_category=&amp;_crumb=dc26c18ae1af65dbdec3f7487d0fbb678840a5c55849aa23b4ae636161cf8593" data-rapid_p="414">
                                                                        <img class="Product__imageData" src="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/1198328c2d101a68702a8996f2254cce6da0456c/i-img900x1200-1645951291o3mldo875826.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" alt="☆ポケモンカード☆旧裏☆ピカチュウ " width="75" height="100" loading="lazy">
                                                                                    </a>
                    <div class="HideSeller">
                                            <a class="HideSeller__link js-rapid-override" data-auction-id="s1039795118" data-auction-category="2084259479" href="https://auctions.yahoo.co.jp/search/search?p=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;va=%E3%83%94%E3%82%AB%E3%83%81%E3%83%A5%E3%82%A6&amp;exflg=1&amp;exsid=anonymous4&amp;b=1&amp;n=50&amp;s1=new&amp;o1=d&amp;mode=4" rel="nofollow" data-ylk="rsec:aal;slk:dslc;pos:5;arbn:" data-rapid_p="415">非表示
                            <div class="HideSeller__balloon">
                                <div class="HideSeller__detail">
                                    <p class="HideSeller__message">この出品者の商品を非表示にする</p>
                                </div>
                            </div>
                        </a>
                                        </div><!-- ./HideSeller -->
                                            <span class="Product__time">1日</span>
                                                                                            <div class="Product__button js-watch-button notlogin">
                                <a class="Button Button--watch is-off rapid-noclick-resp cl-noclick-log js-watch-sync" rel="nofollow" href="#dummy" data-deleteurl="https://auctions.yahoo.co.jp/watchlist/delete?aID=s1039795118&amp;.crumb=" data-addurl="https://auctions.yahoo.co.jp/watchlist/add?aID=s1039795118&amp;.crumb=" data-loading="0" data-operation="add" data-aid="s1039795118" data-ylk="rsec:aal;slk:wc;pos:5;atax:0;catid:2084259479;cid:s1039795118;st:1645951292;end:1646059292;prat:0;spmwl:0;op:;arbn:;sw:off;grat:99.4;ppstr:;cprate:;seltyp:0;best:" data-rapid_p="416">ウォッチ</a>
                                                                </div>
                                                <div class="Layer js-watch-button-layer"></div>
                                    </div><!-- /.Product__image -->
                <div class="Product__detail">
                    <div class="Product__bonus " data-auction-id="s1039795118" data-auction-endtime="1646059292" data-auction-buynowprice="0" data-auction-categoryidpath="0,25464,27727,25826,2084241343,2084259479" data-auction-caneasypayment="1" data-auction-price="500" data-auction-sellerid="anonymous4" data-auction-startprice="500" data-auction-isshoppingitem=""></div>
                    <h3 class="Product__title">
                        <a class="Product__titleLink js-rapid-override js-browseHistory-add" data-auction-id="s1039795118" data-auction-category="2084259479" data-auction-title="☆ポケモンカード☆旧裏☆ピカチュウ " data-auction-img="https://auc-pctr.c.yimg.jp/i/auctions.c.yimg.jp/images.auctions.yahoo.co.jp/image/dr000/auc0302/users/1198328c2d101a68702a8996f2254cce6da0456c/i-img900x1200-1645951291o3mldo875826.jpg?pri=l&amp;w=300&amp;h=300&amp;up=0&amp;nf_src=sy&amp;nf_path=images/auc/pc/top/image/1.0.3/na_170x170.png&amp;nf_st=200" data-auction-price="500" data-auction-isfreeshipping="" href="https://page.auctions.yahoo.co.jp/jp/auction/s1039795118" target="_blank" rel="noopener" data-ylk="rsec:aal;slk:tc;pos:5;atax:0;arbn:;catid:2084259479;cid:s1039795118;etc:p=500,etm=1646059292,stm=1645951292;st:1645951292;end:1646059292;prat:0;op:;grat:99.4;ppstr:;cprate:;seltyp:0;best:" title="☆ポケモンカード☆旧裏☆ピカチュウ " data-rapid_p="417">☆ポケモンカード☆旧裏☆ピカチュウ </a>
                    </h3>
                     <span class="Product__price">
                        <span class="Product__priceValue u-textRed">500円</span>
                    </span><!-- ./Product__price -->
                                            <span class="Product__price">
                            <span class="Product__label">即決</span>
                                                            <span class="Product__priceValue">-</span>
                                                    </span><!-- ./Product__price -->
                                                        </div><!-- /.Product__detail -->
            </li>
                </ul>
    </div>
        <div class="Immersive js-immersive-window" id="superGridContents">
        <div class="Immersive__inner js-immersive-inner">
            <header class="Immersive__header">
                <button class="Immersive__close js-immersive-close" data-ylk="rsec:clos;slk:off;pos:1" data-rapid_p="798">✕</button>
                <div class="Immersive__controllers">
                    <a href="#dummy" class="Immersive__controller Immersive__controller--prev js-immersive-prev" data-ylk="rsec:nextitm;slk:prev;pos:1" data-rapid_p="799">前の商品</a>
                    <a href="#dummy" class="Immersive__controller Immersive__controller--next js-immersive-next" data-ylk="rsec:nextitm;slk:next;pos:1" data-rapid_p="800">次の商品</a>
                </div>
            </header>
            <div class="Immersive__body js-immersive-body"></div>
            <div class="Immersive__footer  js-immersive-footer"></div>
        </div>
    </div>
</div>
`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Errorf("could not create document: %v", err)
	}

	got := make([]Listing, 0)
	doc.Find("li.Product").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		listing, _ := s.parseItem(sel, nil, "")
		if listing != nil {
			got = append(got, *listing)
		}
		return true
	})

	exp := []Listing{
		{
			URL:                 "https://page.auctions.yahoo.co.jp/jp/auction/s1039793076",
			Title:               "【ポケモンカード】200/SM-P PROMO ピカチュウ プロモ　中古　ポケカ",
			PriceY:              "500",
			PriceCurrency:       "4",
			BuyNowPriceY:        "0",
			BuyNowPriceCurrency: "0",
			ID:                  "s1039793076",
			Endtime:             "2022-03-03 13:53:24",
			SellerID:            "anonymous1",
		},
		{
			URL:                 "https://page.auctions.yahoo.co.jp/jp/auction/s1039791085",
			Title:               "ポケモンカード　ボスの指令　(s9)(s4a)(s8b) フュージョンアーツ　VMAX スターバース　ピカチュウ　ミュウ　リザードン",
			PriceY:              "1000",
			PriceCurrency:       "9",
			BuyNowPriceY:        "0",
			BuyNowPriceCurrency: "0",
			ID:                  "s1039791085",
			Endtime:             "2022-03-06 11:51:58",
			SellerID:            "anonymous2",
		},
		{
			URL:                 "https://page.auctions.yahoo.co.jp/jp/auction/r1039797461",
			Title:               "ぬいぐるみ Pokmon Cafe パティシエピカチュウ",
			PriceY:              "5278",
			PriceCurrency:       "46",
			BuyNowPriceY:        "5278",
			BuyNowPriceCurrency: "46",
			ID:                  "r1039797461",
			Endtime:             "2022-02-28 15:45:55",
			SellerID:            "anonymous3",
		},
		{
			URL:                 "https://page.auctions.yahoo.co.jp/jp/auction/x1039785922",
			Title:               "【ポケモンカード】283/SM-P ヨコハマのピカチュウ プロモ　中古　ポケカ",
			PriceY:              "1000",
			PriceCurrency:       "9",
			BuyNowPriceY:        "0",
			BuyNowPriceCurrency: "0",
			ID:                  "x1039785922",
			Endtime:             "2022-03-04 14:42:59",
			SellerID:            "anonymous1",
		},
		{
			URL:                 "https://page.auctions.yahoo.co.jp/jp/auction/s1039795118",
			Title:               "☆ポケモンカード☆旧裏☆ピカチュウ ",
			PriceY:              "500",
			PriceCurrency:       "4",
			BuyNowPriceY:        "0",
			BuyNowPriceCurrency: "0",
			ID:                  "s1039795118",
			Endtime:             "2022-02-28 15:41:32",
			SellerID:            "anonymous4",
		},
	}

	if !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v but got %+v", exp, got)
	}
}
