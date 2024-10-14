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
const newGameButton = document.getElementById("new-game-button");
newGameButton.addEventListener("click", () => startNewGame());
function startNewGame() {
    return __awaiter(this, void 0, void 0, function* () {
        console.log("STARTING NEW GAME");
        try {
            const response = yield fetch("new-game", {
                method: "POST",
            });
            if (!response.ok) {
                throw new Error("Could not fetch new chess game");
            }
            const data = (yield response.json());
            console.log("Fetched game state:", data);
            renderBoard(data);
        }
        catch (error) {
            console.error("There was a problem with fetching a new game:", error);
        }
    });
}
const gameContainer = document.getElementById("game-container");
function renderBoard(gameState) {
    gameContainer.innerHTML = "";
    const boardDiv = document.createElement("div");
    gameState.board.forEach((row, rowIndex) => {
        const rowDiv = document.createElement("div");
        rowDiv.classList.add("board-row");
        row.forEach((piece, pieceIndex) => {
            const squareDiv = document.createElement("div");
            squareDiv.classList.add("chess-square");
            if ((rowIndex + pieceIndex) & 1) {
                squareDiv.classList.add("light-chess-square");
            }
            else {
                squareDiv.classList.add("dark-chess-square");
            }
            const pieceDiv = document.createElement("div");
            pieceDiv.classList.add("chess-piece");
            pieceDiv.innerText = piece ? piece.symbol : "";
            squareDiv.appendChild(pieceDiv);
            rowDiv.appendChild(squareDiv);
        });
        boardDiv.appendChild(rowDiv);
    });
    gameContainer.appendChild(boardDiv);
}
