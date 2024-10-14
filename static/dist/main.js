"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
document.addEventListener("DOMContentLoaded", () => {
    setupNewGameButton();
});
function setupNewGameButton() {
    const newGameButton = document.getElementById("new-game-button");
    newGameButton.addEventListener("click", startNewGame);
}
function startNewGame() {
    return __awaiter(this, void 0, void 0, function* () {
        console.log("STARTING NEW GAME");
        try {
            const gameState = yield fetchNewGame();
            console.log("Fetched game state:", gameState);
            renderBoard(gameState);
        }
        catch (error) {
            console.error("There was a problem with fetching a new game:", error);
        }
    });
}
function fetchNewGame() {
    return __awaiter(this, void 0, void 0, function* () {
        const response = yield fetch("new-game", { method: "POST" });
        if (!response.ok) {
            throw new Error("Could not fetch new chess game");
        }
        return response.json();
    });
}
const gameContainer = document.getElementById("game-container");
function renderBoard(gameState) {
    gameContainer.innerHTML = "";
    const boardDiv = createBoardDiv(gameState);
    gameContainer.appendChild(boardDiv);
}
function createBoardDiv(gameState) {
    const boardDiv = document.createElement("div");
    gameState.board.forEach((row, rowIndex) => {
        const rowDiv = createRowDiv(row, rowIndex);
        boardDiv.appendChild(rowDiv);
    });
    return boardDiv;
}
function createRowDiv(row, rowIndex) {
    const rowDiv = document.createElement("div");
    rowDiv.classList.add("board-row");
    row.forEach((piece, pieceIndex) => {
        const squareDiv = createSquareDiv(piece, rowIndex, pieceIndex);
        rowDiv.appendChild(squareDiv);
    });
    return rowDiv;
}
function createSquareDiv(piece, rowIndex, pieceIndex) {
    const squareDiv = document.createElement("div");
    squareDiv.classList.add("chess-square");
    if ((rowIndex + pieceIndex) % 2 === 0) {
        squareDiv.classList.add("dark-chess-square");
    }
    else {
        squareDiv.classList.add("light-chess-square");
    }
    squareDiv.addEventListener("click", () => {
        console.log("Square clicked:", { row: rowIndex, col: pieceIndex });
        squareDiv.classList.add("selected-square");
        handleSquareClick({ row: rowIndex, col: pieceIndex });
    });
    const pieceDiv = createPieceDiv(piece);
    squareDiv.appendChild(pieceDiv);
    return squareDiv;
}
function createPieceDiv(piece) {
    const pieceDiv = document.createElement("div");
    pieceDiv.classList.add("chess-piece");
    pieceDiv.innerText = piece ? piece.symbol : "";
    return pieceDiv;
}
let selectedSquares = {};
function handleSquareClick(square) {
    if (!selectedSquares.from) {
        selectedSquares.from = square;
    }
    else {
        selectedSquares.to = square;
        sendMoveRequest();
    }
}
function sendMoveRequest() {
    return __awaiter(this, void 0, void 0, function* () {
        if (selectedSquares.from && selectedSquares.to) {
            const moveRequest = {
                fromX: selectedSquares.from.row,
                fromY: selectedSquares.from.col,
                toX: selectedSquares.to.row,
                toY: selectedSquares.to.col,
            };
            try {
                yield makeMove(moveRequest);
                const gameState = yield currentGameState();
                renderBoard(gameState);
            }
            catch (error) {
                console.error(error);
            }
            finally {
                selectedSquares = {};
            }
        }
    });
}
function makeMove(move) {
    return __awaiter(this, void 0, void 0, function* () {
        const response = yield fetch("/move", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(move),
        });
        if (!response.ok) {
            throw new Error("Move failed");
        }
        return response.json();
    });
}
function currentGameState() {
    return __awaiter(this, void 0, void 0, function* () {
        const response = yield fetch("/current-state", { method: "GET" });
        if (!response.ok) {
            throw new Error("Could not fetch game state");
        }
        return response.json();
    });
}
