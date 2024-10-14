const newGameButton = document.getElementById("new-game-button")!;

newGameButton.addEventListener("click", () => startNewGame());

async function startNewGame() {
  console.log("STARTING NEW GAME");

  try {
    const response = await fetch("new-game", {
      method: "POST",
    });

    if (!response.ok) {
      throw new Error("Could not fetch new chess game");
    }
    const data = (await response.json()) as Game;

    console.log("Fetched game state:", data);

    renderBoard(data);
  } catch (error) {
    console.error("There was a problem with fetching a new game:", error);
  }
}

const gameContainer = document.getElementById("game-container")!;

function renderBoard(gameState: Game) {
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
      } else {
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
