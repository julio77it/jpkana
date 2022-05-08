
var url = {
    'hiragana' : 'https://raw.githubusercontent.com/julio77it/jpkana/main/resources/hiragana.json'
,   'katakana' : 'https://raw.githubusercontent.com/julio77it/jpkana/main/resources/katakana.json'
};
var kanas = {}

async function load(kana) {
    u = url[kana];
    const response = await fetch(u);
    const j = await response.json();
    kanas[kana] = j;
}

load('hiragana')
//load('katakana')

function generate(kana, level, length) {
    console.log(kana, level, length)
    filtered = kanas[kana].filter(function (value) {
        return level >= value['difficulty'];
    });

    toread = ''
    towrite = ''

    for (let i = 0; i < length; i++) {
        idx = Math.floor(Math.random() * filtered.length);
        toread += filtered[idx]['kana']
        towrite += filtered[idx]['romanji']
    }
    return [toread,towrite]
}

