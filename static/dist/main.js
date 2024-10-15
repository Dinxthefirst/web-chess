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
let selectedSquares = {};
function handleSquareClick(square) {
    if (!selectedSquares.from) {
        selectedSquares.from = square;
    }
    else {
        selectedSquares.to = square;
        const move = {
            fromRow: selectedSquares.from.row,
            fromCol: selectedSquares.from.col,
            toRow: selectedSquares.to.row,
            toCol: selectedSquares.to.col,
        };
        sendMoveRequest(move);
        selectedSquares = {};
    }
}
//// ASYNC FUNCTIONS
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
