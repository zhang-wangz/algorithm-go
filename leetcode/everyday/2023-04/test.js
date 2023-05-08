// ==UserScript==
// @name         hsrnsw 小说
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  try to take over the world!
// @author       You
// @match        http://192.168.8.1/*
// @match        https://xiaolincoding.com/*
// @run-at       document-end
// @icon         https://www.google.com/s2/favicons?sz=64&domain=hsrnsw.org
// @require      https://cdn.bootcdn.net/ajax/libs/jquery/3.5.1/jquery.min.js
// @require  https://gist.github.com/raw/2625891/waitForKeyElements.js
// @grant        none
// ==/UserScript==

(function() {
    'use strict';
    let url = "https://www.hsrnsw.org/xshs/103349/22531292.html"
    window.onerror = function(message, url, lineno, colno, error) {
        console.log(message, url, lineno, colno, error)
        return true
    }
    // document.body.style.backgroundColor = '#fafafa'
    // document.querySelector('.box_con').style.border = '2px solid #fafafa'
    $.ajaxPrefilter(function (options) {
        if (options.crossDomain && jQuery.support.cors) {
            let http = (window.location.protocol === 'http:' ? 'http:' : 'https:');
            options.url =  http + '//cors-anywhere.herokuapp.com/' + options.url;
        }
    });

    function testget () {
        let a = document.querySelector('#content')
        let arr = a.childNodes
        let len = arr.length
        let cnt=0
        let cnt1=0
        for (let i = len-1; i >= 0; i-- ) {
            if (arr[i].tagName == 'BR' && cnt <= 66) { arr[i].remove(); cnt++; }
            else if (arr[i].tagName == undefined && cnt1 <= 32) {arr[i].remove(); cnt1++}
        }
        let up = document.querySelector("#wrapper > div.content_read > div > div.bottem2 > a:nth-child(2)")
        let down = document.querySelector("#wrapper > div.content_read > div > div.bottem2 > a:nth-child(4)")
        let upurl = 'https://www.hsrnsw.org/' + up.getAttribute('href')
        let downurl =  'https://www.hsrnsw.org/' + down.getAttribute('href')
        for (const element of arr) {
            if(element.tagName !='BR') {
                console.log(element)
            }
        }
        console.log(document.querySelector('.bookname').childNodes[1].textContent)
        console.log(upurl)
        console.log(downurl)
        return downurl
    }
    let inputU = document.createElement('input')
    inputU.setAttribute('class', "specinput")
    inputU.setAttribute('value', url)
    let button = document.createElement('button')
    button.setAttribute('class', "specbutton")
    button.textContent = 'get'
    let button1 = document.createElement('button')
    button1.setAttribute('class', "specbutton1")
    button1.textContent = 'next'
    if (window.location.href.match("https://xiaolincoding.com/*")) {
        waitForKeyElements("#app > div.theme-container > div:nth-child(4) > main > div:nth-child(1)", ()=>{
            let nav = document.querySelector("#app > div.theme-container > div:nth-child(4) > main > div:nth-child(1)")
            nav.appendChild(inputU)
            nav.appendChild(button)
            nav.appendChild(button1)
        });
    }
    button.onclick = function() {
        console.log($('.specinput').val())
        $.get(
            $('.specinput').val(),
            function (response) {
                let test = document.createElement("div")
                test.setAttribute('hidden', true)
                test.setAttribute('class', 'speccontent')
                document.body.appendChild(test)
                $('.speccontent').html(response)
                let nex = testget()
                button1.onclick = () => {
                    $('.specinput').val(nex)
                    console.log("已切换到下一章...")
                }
            });
    }
})();