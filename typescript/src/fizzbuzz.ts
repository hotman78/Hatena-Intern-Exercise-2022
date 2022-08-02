const CreateFizzBuzzFunc=()=>{
  let count=0
  return ()=>{
    count++
    if(count%15==0)return "FizzBuzz";
    if(count%3==0)return "Fizz";
    if(count%5==0)return "Buzz";
    return String(count);
  }
}

export function initializeFizzBuzz(
  list: HTMLOListElement,
  button: HTMLButtonElement
): void {
  console.debug("initializeFizzBuzz", list, button);
  const FizzBuzzFunc=CreateFizzBuzzFunc();
  button.onclick=()=>{
    const liLast = document.createElement('li');
    liLast.textContent = FizzBuzzFunc();
    list.appendChild(liLast);
  }
}
