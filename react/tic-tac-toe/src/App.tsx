import React, { useState } from 'react';

interface SquareProps {
    value: string;
    onSquareClick: () => void;
}

interface BoardProps {
  xIsNext: boolean;
  squares: Array<string>;
  onPlay: (arr: Array<string>) => void;
}


// functions
const Square: React.FC<SquareProps> = ({ value, onSquareClick }) => {
    return (
        <button 
            className="square"
            onClick={onSquareClick}
        >
            {value}
        </button>
    );
}

function calculateWinner(squares: Array<string>) {
    const lines = [
        [0, 1, 2],
        [3, 4, 5],
        [6, 7, 8],
        [0, 3, 6],
        [1, 4, 7],
        [2, 5, 8],
        [0, 4, 8],
        [2, 4, 6]
    ]

    for(let i = 0; i < lines.length; i++) {
        // number-array 解包
        const [a, b, c] = lines[i];
        if(squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
            return squares[a];
        }
    }

    return null;
}

function Board({ xIsNext, squares, onPlay }: BoardProps): React.JSX.Element {
    // const [squares, setSquares] = useState<Array<string>>(Array(9).fill(null));
    // const [currPattern, setCurrPattern] = useState<string>("X");

    function handleClick(i: number) {
        // closure: squares
        if(squares[i] || calculateWinner(squares)) { return; }

        // .slice(): copy the original array
        const nextSquares = squares.slice();
        if (xIsNext) {
          nextSquares[i] = "X";
        } else {
          nextSquares[i] = "O";
        }
        onPlay(nextSquares);
        // nextSquares[i] = currPattern;
        // // change to next player
        // if(currPattern === "X") { 
        //   setCurrPattern("O"); 
        // } else { 
        //   setCurrPattern("X"); 
        // }
        // setSquares(nextSquares);
    }

    // squares 是Board Component的一个state，由于state改变
    // 导致重新渲染，此处的逻辑也会重新执行
    const winner = calculateWinner(squares);
    let status;
    if (winner) {
      status = "Winner: " + winner;
    } else {
      let next: string = xIsNext ? "X" : "O";
      status = "Next player: " + next;
    }

    return <>
      <div className="status">{status}</div>
      <div className="board-row">
        {/* 此处不可直接调用: handleClick(1)，会陷入render死循环；以箭头函数的方式把callable-obj送进去 */}
        <Square value={squares[0]} onSquareClick={() => handleClick(0)} />
        <Square value={squares[1]} onSquareClick={() => handleClick(1)} />
        <Square value={squares[2]} onSquareClick={() => handleClick(2)} />
      </div>
      <div className="board-row">
        <Square value={squares[3]} onSquareClick={() => handleClick(3)} />
        <Square value={squares[4]} onSquareClick={() => handleClick(4)} />
        <Square value={squares[5]} onSquareClick={() => handleClick(5)} />
      </div>
      <div className="board-row">
        <Square value={squares[6]} onSquareClick={() => handleClick(6)} />
        <Square value={squares[7]} onSquareClick={() => handleClick(7)} />
        <Square value={squares[8]} onSquareClick={() => handleClick(8)} />
      </div>
    </>;
}

// Export Component
export default function Game(): React.JSX.Element {
  // states
  // const [xIsNext, setXIsNext] = useState(true);
  const [history, setHistory] = useState([Array(9).fill(null)]);
  const [currentMove, setCurrentMove] = useState(0);

  // refer from states
  const xIsNext = currentMove % 2 === 0;
  const currentSquares = history[currentMove];

  // Update state function:
  // 经由props传递给board调用，修改state使Game组件重新渲染
  function handlePlay(nextSquares: Array<string>) {
    const nextHistory = [...history.slice(0, currentMove + 1), nextSquares];
    setHistory(nextHistory);
    setCurrentMove(nextHistory.length - 1);
    // setXIsNext(!xIsNext);
  }


  function jumpTo(nextMove: number) {
    setCurrentMove(nextMove);
    // setXIsNext(nextMove % 2 == 0);
  }

  // Game 每次重新渲染时都会执行，用于在右侧显示历史执行记录
  const moves: React.JSX.Element[] = history.map((squares: Array<string>, index: number) => {
    let description;
    if (index > 0) {
      description = 'Go to move #' + index;
    } else {
      description = 'Go to start';
    }

    return (
      // React 每次渲染 list 时需要能确定是哪些元素发生了修改，
      // 因此此处必须加上一个key作为id使用
      <li key={index}>
        <button onClick={() => jumpTo(index)}>{description}</button>;
      </li>
    );
  })

  return (
    <div className="game">
      <div className="game-board">
        <Board xIsNext={xIsNext} squares={currentSquares} onPlay={handlePlay}/>
      </div>
      <div className="game-info">
        <ol>{moves}</ol>
      </div>
    </div>
  )
}