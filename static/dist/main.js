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
// import * as Types from "./types";
// TYPES
const boardsize = 8;
var PieceType;
(function (PieceType) {
    PieceType[PieceType["None"] = 0] = "None";
    PieceType[PieceType["King"] = 1] = "King";
    PieceType[PieceType["Pawn"] = 2] = "Pawn";
    PieceType[PieceType["Knight"] = 3] = "Knight";
    PieceType[PieceType["Bishop"] = 4] = "Bishop";
    PieceType[PieceType["Rook"] = 5] = "Rook";
    PieceType[PieceType["Queen"] = 6] = "Queen";
    PieceType[PieceType["White"] = 8] = "White";
    PieceType[PieceType["Black"] = 16] = "Black";
})(PieceType || (PieceType = {}));
const PieceImages = {
    [PieceType.Pawn]: {
        [PieceType.White]: "/static/images/white_pawn.svg",
        [PieceType.Black]: "/static/images/black_pawn.svg",
    },
    [PieceType.Rook]: {
        [PieceType.White]: "/static/images/white_rook.svg",
        [PieceType.Black]: "/static/images/black_rook.svg",
    },
    [PieceType.Knight]: {
        [PieceType.White]: "/static/images/white_knight.svg",
        [PieceType.Black]: "/static/images/black_knight.svg",
    },
    [PieceType.Bishop]: {
        [PieceType.White]: "/static/images/white_bishop.svg",
        [PieceType.Black]: "/static/images/black_bishop.svg",
    },
    [PieceType.Queen]: {
        [PieceType.White]: "/static/images/white_queen.svg",
        [PieceType.Black]: "/static/images/black_queen.svg",
    },
    [PieceType.King]: {
        [PieceType.White]: "/static/images/white_king.svg",
        [PieceType.Black]: "/static/images/black_king.svg",
    },
};
///// EVERYHING ELSE
document.addEventListener("DOMContentLoaded", () => {
    setupNewGameButton();
});
function setupNewGameButton() {
    const newGameButton = document.getElementById("new-game-button");
    newGameButton.addEventListener("click", startNewGame);
}
const gameContainer = document.getElementById("game-container");
function renderBoard(gameState) {
    gameContainer.innerHTML = "";
    const boardDiv = createBoardDiv(gameState);
    gameContainer.appendChild(boardDiv);
}
function createBoardDiv(gameState) {
    const boardDiv = document.createElement("div");
    for (let rank = 7; rank >= 0; rank--) {
        const rowDiv = document.createElement("div");
        rowDiv.classList.add("board-row");
        for (let file = 0; file < boardsize; file++) {
            const index = rank * boardsize + file;
            const piece = gameState.board[index];
            const squareDiv = createSquareDiv(piece, index);
            squareDiv.dataset.index = index.toString();
            if ((rank + file) % 2 === 0) {
                squareDiv.classList.add("light-chess-square");
            }
            else {
                squareDiv.classList.add("dark-chess-square");
            }
            rowDiv.appendChild(squareDiv);
        }
        boardDiv.appendChild(rowDiv);
    }
    return boardDiv;
}
function createSquareDiv(piece, index) {
    const squareDiv = document.createElement("div");
    squareDiv.classList.add("chess-square");
    squareDiv.textContent = index.toString();
    squareDiv.addEventListener("click", () => {
        handleSquareClick(index);
    });
    const pieceDiv = createPieceDiv(piece);
    squareDiv.appendChild(pieceDiv);
    return squareDiv;
}
function createPieceDiv(piece) {
    const pieceDiv = document.createElement("div");
    pieceDiv.classList.add("chess-piece");
    const pieceType = piece.type & 7;
    if (piece.type === PieceType.None) {
        return pieceDiv;
    }
    const pieceColor = piece.type & 24;
    const pieceImg = document.createElement("img");
    pieceImg.src = PieceImages[pieceType][pieceColor];
    pieceDiv.appendChild(pieceImg);
    return pieceDiv;
}
let selectedSquares = [];
function handleSquareClick(index) {
    if (selectedSquares.length === 0) {
        selectedSquares[0] = index;
        getLegalMoves(index);
        return;
    }
    if (selectedSquares[0] === index) {
        selectedSquares = [];
        removeHighlight();
        return;
    }
    selectedSquares[1] = index;
    const move = {
        startSquare: selectedSquares[0],
        targetSquare: selectedSquares[1],
    };
    sendMoveRequest(move);
    selectedSquares = [];
    removeHighlight();
}
function highlightLegalMoves(moves) {
    removeHighlight();
    for (let i = 0; i < moves.length; i++) {
        highlightSquare(moves[i].targetSquare);
    }
}
function highlightSquare(index) {
    const square = document.querySelector(`[data-index='${index}']`);
    square.classList.add("highlighted-square");
    const circle = document.createElement("div");
    circle.classList.add("highlight-circle");
    square.appendChild(circle);
}
function removeHighlight() {
    const squares = document.getElementsByClassName("chess-square");
    for (let i = 0; i < squares.length; i++) {
        const square = squares[i];
        square.classList.remove("highlighted-square");
        const existingCircle = square.querySelector(".highlight-circle");
        if (existingCircle) {
            existingCircle.remove();
        }
    }
}
//// ASYNC FUNCTIONS
function startNewGame() {
    return __awaiter(this, void 0, void 0, function* () {
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
function sendMoveRequest(move) {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            yield makeMove(move);
            const gameState = yield currentGameState();
            renderBoard(gameState);
        }
        catch (error) {
            console.error(error);
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
function getLegalMoves(index) {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const moves = yield fetchLegalMoves(index);
            console.log("Legal moves:", moves);
            highlightLegalMoves(moves);
        }
        catch (error) {
            console.error("There was a problem with fetching legal moves:", error);
        }
    });
}
function fetchLegalMoves(index) {
    return __awaiter(this, void 0, void 0, function* () {
        const response = yield fetch(`/legal-moves/${index}`, { method: "GET" });
        if (!response.ok) {
            throw new Error("Could not fetch legal moves");
        }
        return response.json();
    });
}
