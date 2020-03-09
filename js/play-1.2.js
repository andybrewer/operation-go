// start at the top window position after a game reset
document.body.scrollTop = 0;
document.documentElement = 0;

// game settings
let gameConfig = {
    operation_go: {
        editorTitle: (levelNum) => { return `~/operation_go/level_${levelNum}.go`; },
        codepath: (levelNum) => { return `./levels/level_${levelNum}/level_${levelNum}.go`; },
        finalePath: './levels/finale.txt',
        imgpath: (levelNum) => { return `./img/level_${levelNum}.png`; },
        levels: 12,
    }
}

// language settings
let langConfig = {
    go: {
        cmd: (levelNum) => { return `> go run level_${levelNum}.go`; },
        exec: async (levelNum, code) => {
            let results = {
                output: "",
                syntaxErr: "",
                fetchErr: "",
            }

            try {
                let response = await fetch("https://play.golang.org/compile", {
                    method: "POST",
                    headers: { "Content-Type": "text/plain" },
                    body: JSON.stringify({
                        body: code
                    })
                });
                let data = await response.json();
                if (data.Events != null) {
                    results.output = data.Events[0].Message;
                    return results;
                }

                results.syntaxErr = data.Errors.replace(/prog.go:/g, `level_${levelNum}.go:`);
                return results;
            } catch (err) {
                results.fetchErr = err.toString();
                return results;
            }
        },
        parse: {
            editStart: "/* EDIT START */",
            editEnd: "/* EDIT END */",
            hintStart: "// HINT: ",
            hintEnd: "\n",
            introStart: "/* INTRO",
            introEnd: "*/",
            nameStart: "//",
            nameEnd: "\n",
            outputStart: "/* OUTPUT",
            outputEnd: "*/",
        },
        syntax: "ace/mode/golang",
    },
}

// newLevel - track the state of a level
function newLevel() {
    return {
        code: "",
        editorId: "",
        hint: "",
        id: "",
        intro: "",
        name: "",
        output: "",
        readOnlyTop: 0,
        readOnlyBot: 0,
        solution: "",
    }
}

let game = {
    gameId: "",
    langId: "",
    levels: [],
};

// setupGame - initialize a new game or restore a saved one
async function setupGame(gameId, langId) {
    let savedState = localStorage.getItem(gameId);

    // initialize a new game
    if (savedState === null) {
        game.gameId = gameId;
        game.langId = langId;
        game.levels = [newLevel()];
        preload();
        renderLevel(1);
        return;
    }

    // restore a saved game
    game = JSON.parse(savedState);

    preload();

    // render the previous levels
    for (let i = 0; i < game.levels.length; i++) {
        await renderLevel(i + 1);
        if (i + 1 < game.levels.length) {
            // render the code solution
            ace.edit(document.querySelector(`#level_${i + 1} .editor__window`)).setValue(game.levels[i].solution, -1);

            // render the success output
            document.querySelector(`#level_${i + 1} .output__cmd`).innerHTML = langConfig[game.langId].cmd(i + 1);
            updateOutput(i + 1, game.levels[i].output, "Success!", "output__message--pass");

            // lock the level
            lockLevel(i + 1);
        }
    }

    setTimeout(function () {
        document.querySelector(`#level_${game.levels.length}`).scrollIntoView({
            behavior: "smooth",
            block: "start",
            inline: "start",
        });
    }, 500);
}

// preload - perform some load ahead optimization
function preload() {
    var images = [];
    for (var i = 1; i <= gameConfig[game.gameId].levels; i++) {
        var imgPath = gameConfig[game.gameId].imgpath(i);
        var img = new Image();
        img.src = imgPath;
    }
}

// resetGame - resets all saved progress
function resetGame(gameId) {
    if (confirm("Do you want to clear your progress and start over?")) {
        localStorage.removeItem(gameId);
        location.reload();
    }

    return false;
}

// renderLevel - downloads and parses the current level
async function renderLevel(levelNum) {
    let level = game.levels[levelNum - 1];

    // check for victory
    if (levelNum > gameConfig[game.gameId].levels) {
        await fetch(gameConfig[game.gameId].finalePath, {
            method: "GET",
        })
            .then(function (response) {
                return response.text();
            })
            .then(async function (content) {
                // copy our level HTML template and remove elements we don't need
                let levelContent = document.querySelector("#level").content.cloneNode(true);
                levelContent.removeChild(levelContent.querySelector(".editor"));
                levelContent.removeChild(levelContent.querySelector(".cmd"));

                // create a new level node and append it to the DOM
                levelDOM = document.createElement("div");
                levelDOM.id = `level_${levelNum}`;
                await levelDOM.appendChild(levelContent);
                await document.querySelector("#play").appendChild(levelDOM);

                // add the content
                levelDOM.querySelector(".story__title").innerHTML = "Finale";
                levelDOM.querySelector(".story__title").classList += " story__title--finale";
                levelDOM.querySelector(".story__intro").innerHTML = toParagraphs(content.split("\n")) + `<div class="story__reset"><a href="https://www.twitter.com/intent/tweet?text=I%20just%20beat%20Operation%20Go%2C%20an%20online%20%23golang%20game.%20See%20if%20you%20can%20beat%20it%20too%3A%20https://andybrewer.github.io/operation-go/" target="_blank">Tweet your Victory &rarr;</a><br /><br />or <a href="#" onclick="return resetGame('${game.gameId}');">play again</a></div>`;

                // scroll the victory into view
                document.querySelector(`#level_${levelNum}`).scrollIntoView({
                    behavior: "smooth",
                    block: "start",
                    inline: "start",
                });
            })
            .catch(async err => {
                alert("Sorry, there was an error loading the finale. Please try reloading the page and if the issue persists contact us with a screenshot of the issue.");
                console.error(err);
            });
        return;
    }

    await fetch(gameConfig[game.gameId].codepath(levelNum), {
        method: "GET",
    })
        .then(function (response) {
            return response.text();
        })
        .then(async function (content) {
            // copy our level HTML template
            let levelContent = document.querySelector("#level").content.cloneNode(true);

            // create a new level node and append it to the DOM
            levelDOM = document.createElement("div");
            levelDOM.id = `level_${levelNum}`;
            await levelDOM.appendChild(levelContent);
            await document.querySelector("#play").appendChild(levelDOM);

            // render Ace editor
            let levelEditor = ace.edit(document.querySelector(`#level_${levelNum} .editor__window`), {
                fontFamily: "monospace",
                fontSize: "14px",
                maxLines: 1000,
                minLines: 10,
                mode: langConfig[game.langId].syntax,
                selectionStyle: "text",
                showPrintMargin: false,
                wrap: true,
            });

            // re-render read-only lines
            levelEditor.renderer.on("afterRender", function () {
                styleReadOnlyLines(levelNum, level.readOnlyTop, level.readOnlyBot);
            });

            // read-only lines
            levelEditor.commands.on("exec", function (e) {
                if (e.command.readOnly) { return; }

                var rowCol = levelEditor.selection.getCursor();
                if (rowCol.row < level.readOnlyTop || rowCol.row >= levelEditor.session.getLength() - level.readOnlyBot) {
                    e.preventDefault();
                    e.stopPropagation();
                }
            });

            // parse the level
            parseLevel(levelNum, content, game);
        })
        .catch(async err => {
            alert("Sorry, there was an error loading the level. Please try reloading the page and if the issue persists contact us with a screenshot of the issue.");
            console.error(err);
        });
}

// run - runs the code in the current editor and validates the output
async function run() {
    let levelNum = game.levels.length;
    let level = game.levels[levelNum - 1];

    resetOutput(levelNum, langConfig[game.langId].cmd());

    let code = ace.edit(document.querySelector(`#level_${levelNum} .editor__window`)).getValue();

    // clear any previous output and run the new code
    resetOutput(levelNum, langConfig[game.langId].cmd(levelNum));
    let results = await langConfig[game.langId].exec(levelNum, code);
    validateOutput(levelNum, results, code);
}

// parseLevel - extracts level info from the source code
function parseLevel(levelNum, content, game) {
    let parser = langConfig[game.langId].parse;
    let level = game.levels[levelNum - 1];

    // Extract the level name
    let nameStartPos = content.indexOf(parser.nameStart);
    let nameEndPos = content.indexOf(parser.nameEnd, nameStartPos);
    level.name = content.substr(nameStartPos + parser.nameStart.length, nameEndPos - parser.nameStart.length).trim();
    content = content.substr(nameEndPos + parser.nameEnd.length).trim();

    // Extract the intro
    let introStartPos = content.indexOf(parser.introStart);
    let introEndPos = content.indexOf(parser.introEnd, introStartPos);
    level.intro = toParagraphs(content.substr(introStartPos + parser.introStart.length, introEndPos - parser.introStart.length).trim().split("\n"));
    content = content.substr(introEndPos + parser.introEnd.length).trim();

    // Extract the hint
    let hintStartPos = content.indexOf(parser.hintStart);
    let hintEndPos = content.indexOf(parser.hintEnd, hintStartPos);
    level.hint = content.substr(hintStartPos + parser.hintStart.length, hintEndPos - parser.hintStart.length).trim();
    content = content.substr(hintEndPos + parser.hintEnd.length).trim();

    // Extract the output
    let outputStartPos = content.indexOf(parser.outputStart);
    let outputEndPos = content.indexOf(parser.outputEnd, outputStartPos);
    level.output = content.substr(outputStartPos + parser.outputStart.length, outputEndPos - outputStartPos - parser.outputStart.length).trim();
    content = content.substr(0, outputStartPos).trim();

    // Remove edit flags
    let editStartPos = content.indexOf(parser.editStart);
    level.readOnlyTop = content.substr(0, editStartPos).split("\n").length;
    let editEndPos = content.indexOf(parser.editEnd);
    level.readOnlyBot = content.substr(editEndPos).split("\n").length;
    level.code = content.substr(0, editStartPos).trim() + "\n" + content.substr(editStartPos + parser.editStart.length, editEndPos - editStartPos - parser.editStart.length) + content.substr(editEndPos + parser.editEnd.length);

    // update the content
    document.querySelector(`#level_${levelNum} .story__pic`).src = gameConfig[game.gameId].imgpath(levelNum);
    document.querySelector(`#level_${levelNum} .story__title`).innerHTML = level.name;
    document.querySelector(`#level_${levelNum} .story__intro`).innerHTML = level.intro;
    document.querySelector(`#level_${levelNum} .editor__nav--title`).innerHTML = gameConfig[game.gameId].editorTitle(levelNum);

    // render the code in the editor
    ace.edit(document.querySelector(`#level_${levelNum} .editor__window`)).setValue(level.code, -1);

    // move the window so the new level can be seen
    if (levelNum > 1 && levelNum == game.levels.length) {
        document.querySelector(`#level_${levelNum}`).scrollIntoView({
            behavior: "smooth",
            block: "start",
            inline: "start",
        });
    }
}

// toParagraphs - convert an array to paragraphs
function toParagraphs(arr) {
    return arr.filter(el => { return el != "" }).map(el => { return `<p>${el}</p>` }).join("");
}

// resetOutput - clear any previous output
function resetOutput(levelNum, cmdString) {
    document.querySelector(`#level_${levelNum} .output__cmd`).innerHTML = cmdString;
    document.querySelector(`#level_${levelNum} .output__results`).innerHTML = "";
    document.querySelector(`#level_${levelNum} .output__message`).innerHTML = "";
    document.querySelector(`#level_${levelNum} .output__message`).classList = "output__message";
}

// validateOutput - checks if the run output against the level's expected output
function validateOutput(levelNum, results, code) {
    let level = game.levels[levelNum - 1];

    // check if we got a server error
    if (results === undefined) {
        updateOutput(levelNum, "results undefined", "Sorry, there was a server error.", "output__message--fail");
        return;
    }

    // check if we got a server error
    if (results.fetchErr != "") {
        updateOutput(levelNum, results.fetchErr, "Sorry, there was a server error.", "output__message--fail");
        return;
    }

    // check if we got a syntax error
    if (results.syntaxErr != "") {
        updateOutput(levelNum, results.syntaxErr, "Sorry, there was a syntax error.", "output__message--fail");
        return;
    }

    // check if we got the correct output
    if (results.output.trim() != level.output) {
        updateOutput(levelNum, results.output, "Sorry, that's not correct", "output__message--fail");
        return;
    }

    // confirm success
    updateOutput(levelNum, results.output, "Success!", "output__message--pass");

    // save the progress
    level.solution = code;
    game.levels.push(newLevel());
    localStorage.setItem(game.gameId, JSON.stringify(game));

    // load the next level
    setTimeout(() => {
        lockLevel(levelNum);
        renderLevel(levelNum + 1);
    }, 1500);
}

// lockLevel - makes a level uneditable
function lockLevel(levelNum) {
    styleReadOnlyLines(levelNum, 9999, 9999);

    ace.edit(document.querySelector(`#level_${levelNum} .editor__window`)).renderer.on("afterRender", function () {
        setTimeout(function () {
            styleReadOnlyLines(levelNum, 9999, 9999);
        }, 0);
    });

    ace.edit(document.querySelector(`#level_${levelNum} .editor__window`)).commands.on("exec", function (e) {
        if (e.command.readOnly) { return; }

        e.preventDefault();
        e.stopPropagation();
    });

    document.querySelector(`#level_${levelNum} .editor__nav--hint`).style = "display: none;";
    document.querySelector(`#level_${levelNum} .cmd`).style = "display: none;";
}

// updateOutput - renders the output to the screen
function updateOutput(levelNum, results, message, messageClass) {
    results = results.replace(/(\r\n|\n|\r)/g, "<br />");

    document.querySelector(`#level_${levelNum} .output__results`).innerHTML = results;
    document.querySelector(`#level_${levelNum} .output__message`).innerHTML = message;
    document.querySelector(`#level_${levelNum} .output__message`).classList.add(messageClass);

    // move the window so the results can be seen
    document.querySelector(`#level_${levelNum} .cmd`).scrollIntoView({
        behavior: "smooth",
        block: "nearest",
        inline: "start",
    });
}

// resetCode - restores the original code for the last level
/*
function resetCode() {
    let levelNum = game.levels.length;
    let level = game.levels[levelNum - 1];
    ace.edit(document.querySelector(`#level_${levelNum} .editor__window`)).setValue(level.code, -1);
    resetOutput(levelNum, ">");
}
*/

// showHint - shows the code hint for the level
function showHint() {
    alert(game.levels[game.levels.length - 1].hint);
}

/*
    Ace styling
*/

// style read-only lines
function styleReadOnlyLines(levelNum, top, bottom) {
    let lines = document.querySelectorAll(`#level_${levelNum} .ace_line_group`);

    for (let i = 0; i < lines.length; i++) {
        if (i < top || i >= lines.length - bottom) {
            lines[i].style.opacity = "0.7";
            lines[i].style.background = "#eee";
        } else {
            lines[i].style.opacity = "1";
            lines[i].style.background = "";
        }
    }

    let gutterCells = document.querySelectorAll(`#level_${levelNum} .ace_gutter-cell`);

    for (let i = 0; i < gutterCells.length; i++) {
        if (i < top || i >= gutterCells.length - bottom) {
            gutterCells[i].style.opacity = "0.3";
        } else {
            gutterCells[i].style.opacity = "1";
        }
    }
}