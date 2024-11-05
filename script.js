// katakana should be kanaList

const kanaList = [
    { kana: "あ", romaji: "a", audio: "a.mp3" },
    { kana: "い", romaji: "i", audio: "i.mp3" },
    { kana: "う", romaji: "u", audio: "u.mp3" },
    { kana: "え", romaji: "e", audio: "e.mp3" },
    { kana: "お", romaji: "o", audio: "o.mp3" },
    { kana: "か", romaji: "ka", audio: "ka.mp3" },
    { kana: "き", romaji: "ki", audio: "ki.mp3" },
    { kana: "く", romaji: "ku", audio: "ku.mp3" },
    { kana: "け", romaji: "ke", audio: "ke.mp3" },
    { kana: "こ", romaji: "ko", audio: "ko.mp3" },
    { kana: "さ", romaji: "sa", audio: "sa.mp3" },
    { kana: "し", romaji: "shi", audio: "shi.mp3" },
    { kana: "す", romaji: "su", audio: "su.mp3" },
    { kana: "せ", romaji: "se", audio: "se.mp3" },
    { kana: "そ", romaji: "so", audio: "so.mp3" },
    { kana: "た", romaji: "ta", audio: "ta.mp3" },
    { kana: "ち", romaji: "chi", audio: "chi.mp3" },
    { kana: "つ", romaji: "tsu", audio: "tsu.mp3" },
    { kana: "て", romaji: "te", audio: "te.mp3" },
    { kana: "と", romaji: "to", audio: "to.mp3" },
    { kana: "な", romaji: "na", audio: "na.mp3" },
    { kana: "に", romaji: "ni", audio: "ni.mp3" },
    { kana: "ぬ", romaji: "nu", audio: "nu.mp3" },
    { kana: "ね", romaji: "ne", audio: "ne.mp3" },
    { kana: "の", romaji: "no", audio: "no.mp3" },
    { kana: "は", romaji: "ha", audio: "ha.mp3" },
    { kana: "ひ", romaji: "hi", audio: "hi.mp3" },
    { kana: "ふ", romaji: "fu", audio: "fu.mp3" },
    { kana: "へ", romaji: "he", audio: "he.mp3" },
    { kana: "ほ", romaji: "ho", audio: "ho.mp3" },
    { kana: "ま", romaji: "ma", audio: "ma.mp3" },
    { kana: "み", romaji: "mi", audio: "mi.mp3" },
    { kana: "む", romaji: "mu", audio: "mu.mp3" },
    { kana: "め", romaji: "me", audio: "me.mp3" },
    { kana: "も", romaji: "mo", audio: "mo.mp3" },
    { kana: "や", romaji: "ya", audio: "ya.mp3" },
    { kana: "ゆ", romaji: "yu", audio: "yu.mp3" },
    { kana: "よ", romaji: "yo", audio: "yo.mp3" },
    { kana: "ら", romaji: "ra", audio: "ra.mp3" },
    { kana: "り", romaji: "ri", audio: "ri.mp3" },
    { kana: "る", romaji: "ru", audio: "ru.mp3" },
    { kana: "れ", romaji: "re", audio: "re.mp3" },
    { kana: "ろ", romaji: "ro", audio: "ro.mp3" },
    { kana: "わ", romaji: "wa", audio: "wa.mp3" },
    { kana: "を", romaji: "wo", audio: "wo.mp3" },
    { kana: "ん", romaji: "n", audio: "n.mp3" },
    { kana: "が", romaji: "ga", audio: "ga.mp3" },
    { kana: "ぎ", romaji: "gi", audio: "gi.mp3" },
    { kana: "ぐ", romaji: "gu", audio: "gu.mp3" },
    { kana: "げ", romaji: "ge", audio: "ge.mp3" },
    { kana: "ご", romaji: "go", audio: "go.mp3" },
    { kana: "ざ", romaji: "za", audio: "za.mp3" },
    { kana: "じ", romaji: "ji", audio: "ji.mp3" },
    { kana: "ず", romaji: "zu", audio: "zu.mp3" },
    { kana: "ぜ", romaji: "ze", audio: "ze.mp3" },
    { kana: "ぞ", romaji: "zo", audio: "zo.mp3" },
    { kana: "だ", romaji: "da", audio: "da.mp3" },
    { kana: "ぢ", romaji: "dzi", audio: "dzi.mp3" },
    { kana: "づ", romaji: "dzu", audio: "dzu.mp3" },
    { kana: "で", romaji: "de", audio: "de.mp3" },
    { kana: "ど", romaji: "do", audio: "do.mp3" },
    { kana: "ば", romaji: "ba", audio: "ba.mp3" },
    { kana: "び", romaji: "bi", audio: "bi.mp3" },
    { kana: "ぶ", romaji: "bu", audio: "bu.mp3" },
    { kana: "べ", romaji: "be", audio: "be.mp3" },
    { kana: "ぼ", romaji: "bo", audio: "bo.mp3" },
    { kana: "ぱ", romaji: "pa", audio: "pa.mp3" },
    { kana: "ぴ", romaji: "pi", audio: "pi.mp3" },
    { kana: "ぷ", romaji: "pu", audio: "pu.mp3" },
    { kana: "ぺ", romaji: "pe", audio: "pe.mp3" },
    { kana: "ぽ", romaji: "po", audio: "po.mp3" },
    // Katakana
    { kana: "ア", romaji: "a", audio: "a.mp3" },
    { kana: "イ", romaji: "i", audio: "i.mp3" },
    { kana: "ウ", romaji: "u", audio: "u.mp3" },
    { kana: "エ", romaji: "e", audio: "e.mp3" },
    { kana: "オ", romaji: "o", audio: "o.mp3" },
    { kana: "カ", romaji: "ka", audio: "ka.mp3" },
    { kana: "キ", romaji: "ki", audio: "ki.mp3" },
    { kana: "ク", romaji: "ku", audio: "ku.mp3" },
    { kana: "ケ", romaji: "ke", audio: "ke.mp3" },
    { kana: "コ", romaji: "ko", audio: "ko.mp3" },
    { kana: "サ", romaji: "sa", audio: "sa.mp3" },
    { kana: "シ", romaji: "shi", audio: "shi.mp3" },
    { kana: "ス", romaji: "su", audio: "su.mp3" },
    { kana: "セ", romaji: "se", audio: "se.mp3" },
    { kana: "ソ", romaji: "so", audio: "so.mp3" },
    { kana: "タ", romaji: "ta", audio: "ta.mp3" },
    { kana: "チ", romaji: "chi", audio: "chi.mp3" },
    { kana: "ツ", romaji: "tsu", audio: "tsu.mp3" },
    { kana: "テ", romaji: "te", audio: "te.mp3" },
    { kana: "ト", romaji: "to", audio: "to.mp3" },
    { kana: "ナ", romaji: "na", audio: "na.mp3" },
    { kana: "ニ", romaji: "ni", audio: "ni.mp3" },
    { kana: "ヌ", romaji: "nu", audio: "nu.mp3" },
    { kana: "ネ", romaji: "ne", audio: "ne.mp3" },
    { kana: "ノ", romaji: "no", audio: "no.mp3" },
    { kana: "ハ", romaji: "ha", audio: "ha.mp3" },
    { kana: "ヒ", romaji: "hi", audio: "hi.mp3" },
    { kana: "フ", romaji: "fu", audio: "fu.mp3" },
    { kana: "ヘ", romaji: "he", audio: "he.mp3" },
    { kana: "ホ", romaji: "ho", audio: "ho.mp3" },
    // { kana: "マ", romaji: "ma", audio: "ma.mp3" },
    // { kana: "ミ", romaji: "mi", audio: "mi.mp3" },
    // { kana: "ム", romaji: "mu", audio: "mu.mp3" },
    // { kana: "メ", romaji: "me", audio: "me.mp3" },
    // { kana: "モ", romaji: "mo", audio: "mo.mp3" },
    // { kana: "ヤ", romaji: "ya", audio: "ya.mp3" },
    // { kana: "ユ", romaji: "yu", audio: "yu.mp3" },
    // { kana: "ヨ", romaji: "yo", audio: "yo.mp3" },
    // { kana: "ラ", romaji: "ra", audio: "ra.mp3" },
    // { kana: "リ", romaji: "ri", audio: "ri.mp3" },
    // { kana: "ル", romaji: "ru", audio: "ru.mp3" },
    // { kana: "レ", romaji: "re", audio: "re.mp3" },
    // { kana: "ロ", romaji: "ro", audio: "ro.mp3" },
    // { kana: "ワ", romaji: "wa", audio: "wa.mp3" },
    // { kana: "ヲ", romaji: "wo", audio: "wo.mp3" },
    // { kana: "ン", romaji: "n", audio: "n.mp3" }
];

let currentIndex = 0;

function displayKana() {
    const kanaDisplay = document.getElementById("kanaDisplay");
    const romajiDisplay = document.getElementById("romajiDisplay");
    const autoShowRomaji = document.getElementById("autoShowRomaji").checked;

    const kana = kanaList[currentIndex];
    kanaDisplay.textContent = kana.kana;
    romajiDisplay.textContent = kana.romaji;
    romajiDisplay.style.display = autoShowRomaji ? "block" : "none";

    if (document.getElementById("autoPlayVoice").checked) {
        playAudio(kana.audio);
    }
}

function playAudio(audioFile) {
    const audio = new Audio(`audio/${audioFile}`);
    audio.play();
}

function nextKana() {
    currentIndex = (currentIndex + 1) % kanaList.length;
    displayKana();
}

function previousKana() {
    currentIndex = (currentIndex - 1 + kanaList.length) % kanaList.length;
    displayKana();
}

document.addEventListener("keydown", (event) => {
    if (event.key === "ArrowUp") {
        previousKana();
    } else if (event.key === "ArrowDown") {
        nextKana();
    }
});

// Initialize with a random kana
currentIndex = Math.floor(Math.random() * kanaList.length);
displayKana();
