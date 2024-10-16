// import * as Types from "./types";
// TYPES
const boardSize = 8;

type Color = number;

enum PieceType {
  None,
  King,
  Pawn,
  Knight,
  Bishop,
  Rook,
  Queen,
  White = 8,
  Black = 16,
}

interface Piece {
  type: number;
}

interface Game {
  board: Piece[];
  ColorToMove: Color;
}

interface Move {
  startSquare: number;
  targetSquare: number;
}

const PieceImages: { [key in PieceType]?: { [key in Color]: string } } = {
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

///// FUNCTIONS

document.addEventListener("DOMContentLoaded", () => {
  setupNewGameButton();
});

function setupNewGameButton() {
  const newGameButton = document.getElementById("new-game-button")!;
  newGameButton.addEventListener("click", startNewGame);
}

const gameContainer = document.getElementById("game-container")!;

function renderBoard(game: Game) {
  gameContainer.innerHTML = "";
  gameContainer.dataset.colorToMove = game.ColorToMove.toString();
  const boardDiv = createBoardDiv(game);
  gameContainer.appendChild(boardDiv);
}

function createBoardDiv(game: Game): HTMLDivElement {
  const boardDiv = document.createElement("div");

  for (let rank = 7; rank >= 0; rank--) {
    const rowDiv = document.createElement("div");
    rowDiv.classList.add("board-row");

    for (let file = 0; file < boardSize; file++) {
      const index = rank * boardSize + file;
      const piece = game.board[index];

      const squareDiv = createSquareDiv(piece, index);
      squareDiv.dataset.index = index.toString();
      squareDiv.dataset.type = piece.type.toString();

      if ((rank + file) % 2 === 0) {
        squareDiv.classList.add("dark-chess-square");
      } else {
        squareDiv.classList.add("light-chess-square");
      }

      if (rank === 0) {
        const fileLabel = document.createElement("div");
        fileLabel.classList.add("file-label");
        fileLabel.textContent = String.fromCharCode(97 + file);
        squareDiv.appendChild(fileLabel);
      }

      if (file === 7) {
        const rankLabel = document.createElement("div");
        rankLabel.classList.add("rank-label");
        rankLabel.textContent = String(rank + 1);
        squareDiv.appendChild(rankLabel);
      }

      rowDiv.appendChild(squareDiv);
    }

    boardDiv.appendChild(rowDiv);
  }

  return boardDiv;
}

function createSquareDiv(piece: Piece, index: number): HTMLDivElement {
  const squareDiv = document.createElement("div");
  squareDiv.classList.add("chess-square");

  squareDiv.addEventListener("click", () => {
    handleSquareClick(index);
  });

  const pieceDiv = createPieceDiv(piece);
  squareDiv.appendChild(pieceDiv);

  return squareDiv;
}

function createPieceDiv(piece: Piece): HTMLDivElement {
  const pieceDiv = document.createElement("div");
  pieceDiv.classList.add("chess-piece");

  const pieceType: PieceType = piece.type & 7;
  if (piece.type === PieceType.None) {
    return pieceDiv;
  }
  const pieceColor: Color = piece.type & 24;

  const pieceImg = document.createElement("img");
  pieceImg.src = PieceImages[pieceType]![pieceColor];
  pieceDiv.appendChild(pieceImg);

  return pieceDiv;
}

let selectedSquares: number[] = [];

function handleSquareClick(index: number) {
  if (selectedSquares.length === 0) {
    const colorToMove = parseInt(gameContainer.dataset.colorToMove!);
    const piece = document.querySelector(
      `[data-index='${index}']`
    ) as HTMLElement;
    const pieceType = parseInt(piece.dataset.type!);
    const pieceColor = pieceType & 24;
    console.log("Color to move:", colorToMove);
    console.log("Piece color:", pieceColor);
    if (pieceColor !== colorToMove) {
      return;
    }
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
  const move: Move = {
    startSquare: selectedSquares[0],
    targetSquare: selectedSquares[1],
  };
  console.log("Move:", move);
  sendMoveRequest(move);
  selectedSquares = [];
  removeHighlight();
}

function highlightLegalMoves(moves: Move[]) {
  removeHighlight();

  if (moves.length === 0) return;

  const startSquareIndex = moves[0].startSquare;
  const startSquare = document.querySelector(
    `[data-index='${startSquareIndex}']`
  ) as HTMLElement;
  startSquare.classList.add("highlight-square");

  for (let i = 0; i < moves.length; i++) {
    if (i === startSquareIndex) continue;
    highlightSquare(moves[i].targetSquare);
  }
}

function highlightSquare(index: number) {
  const square = document.querySelector(
    `[data-index='${index}']`
  ) as HTMLElement;

  const pieceType = parseInt(square.dataset.type!);
  if (pieceType !== PieceType.None) {
    const borderHighlight = document.createElement("div");
    borderHighlight.classList.add("highlight-border");
    square.appendChild(borderHighlight);
  } else {
    const circle = document.createElement("div");
    circle.classList.add("highlight-circle");
    square.appendChild(circle);
  }
}

function removeHighlight() {
  const squares = document.getElementsByClassName("chess-square");
  for (let i = 0; i < squares.length; i++) {
    const square = squares[i] as HTMLElement;
    square.classList.remove("highlight-square");
    const existingCircle = square.querySelector(".highlight-circle");
    if (existingCircle) {
      existingCircle.remove();
    }
    const existingBorder = square.querySelector(".highlight-border");
    if (existingBorder) {
      existingBorder.remove();
    }
  }
}

//// ASYNC FUNCTIONS

async function startNewGame() {
  try {
    const gameState = await fetchNewGame();
    console.log("Fetched game state:", gameState);
    renderBoard(gameState);
  } catch (error) {
    console.error("There was a problem with fetching a new game:", error);
  }
}

async function fetchNewGame(): Promise<Game> {
  const response = await fetch("new-game", { method: "POST" });

  if (!response.ok) {
    throw new Error("Could not fetch new chess game");
  }

  return response.json() as Promise<Game>;
}

async function sendMoveRequest(move: Move) {
  try {
    await makeMove(move);
    const gameState = await currentGameState();
    renderBoard(gameState);
  } catch (error) {
    console.error(error);
  }
}

async function makeMove(move: Move): Promise<string> {
  const response = await fetch("/move", {
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
}

async function currentGameState(): Promise<Game> {
  const response = await fetch("/current-state", { method: "GET" });

  if (!response.ok) {
    throw new Error("Could not fetch game state");
  }

  return response.json();
}

async function getLegalMoves(index: number) {
  try {
    const moves = await fetchLegalMoves(index);
    highlightLegalMoves(moves);
  } catch (error) {
    console.error("There was a problem with fetching legal moves:", error);
  }
}

async function fetchLegalMoves(index: number): Promise<Move[]> {
  const response = await fetch(`/legal-moves/${index}`, { method: "GET" });

  if (!response.ok) {
    throw new Error("Could not fetch legal moves");
  }

  return response.json();
}
