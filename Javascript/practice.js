//Sample Callback function -- start
function fil(val,callback){
	console.log(val)
	let pp = val * 2
	callback(null,pp)
}
fil(2,function(err,data){
	console.log(data)
})
//Sample Callback function -- end



//Sample Promise Function -- start
function samPromise(val){
	return new Promise((resolve,reject) => {
		if(val > 2){
			resolve(val)
		}
		else {
			reject()
		}
	})
}
let popi = samPromise(3)
.then(result => console.log(result))
.catch(() => console.log('Eror'))
//Sample Promise function --end




//Sample aSync Function --start
async function caller(val){
	return val * 2
}
async function callher(val){
	let new_val = await caller(val)
	console.log(new_val)
}
callher(6)
//Sample aSync Function --end

//async function always returns a promise and can be used like this.....
caller(3).then(result => console.log(result))

//map function
function arroo(sam){
	return sam.map(x => x+9)
}
let ppio = arroo([3,4,5])
console.log(ppio)
//

//forEach
function fe(sam){
	let gop = sam.forEach(function(x){
		console.log(x)
	})
fe([1,2,3])
//

/*The Promise.all() method returns a single Promise that resolves when all of the 
promises passed as an iterable have resolved or when the iterable contains no promises. 
It rejects with the reason of the first promise that rejects.
var promise1 = Promise.resolve(3);
var promise2 = 42;
var promise3 = new Promise(function(resolve, reject) {
  setTimeout(resolve, 100, 'foo');
});

Promise.all([promise1, promise2, promise3]).then(function(values) {
  console.log(values);
});
// expected output: Array [3, 42, "foo"]
//
*/

//-----------------------------------------------------------------

/*
It's pretty straightforward with some simple rules:

    Whenever you create a promise in a then, return it - any promise you don't return will not be waited for outside.
    Whenever you create multiple promises, .all them - that way it waits for all the promises and no error from any of them are silenced.
    Whenever you nest thens, you can typically return in the middle - then chains are usually at most 1 level deep.
    Whenever you perform IO, it should be with a promise - either it should be in a promise or it should use a promise to signal its completion.

And some tips:

    Mapping is better done with .map than with for/push - if you're mapping values with a function, map lets you concisely express the notion of applying actions 
    one by one and aggregating the results.
    Concurrency is better than sequential execution if it's free - it's better to execute things concurrently and wait for them Promise.all than to 
    execute things one after the other - each waiting before the next.

Ok, so let's get started:

var items = [1, 2, 3, 4, 5];
var fn = function asyncMultiplyBy2(v){ // sample async action
    return new Promise(resolve => setTimeout(() => resolve(v * 2), 100));
};
// map over forEach since it returns

var actions = items.map(fn); // run the function over all items

// we now have a promises array and we want to wait for it

var results = Promise.all(actions); // pass array of promises

results.then(data => // or just .then(console.log)
    console.log(data) // [2, 4, 6, 8, 10]
);

// we can nest this of course, as I said, `then` chains:

var res2 = Promise.all([1, 2, 3, 4, 5].map(fn)).then(
    data => Promise.all(data.map(fn))
).then(function(data){
    // the next `then` is executed after the promise has returned from the previous
    // `then` fulfilled, in this case it's an aggregate promise because of 
    // the `.all` 
    return Promise.all(data.map(fn));
}).then(function(data){
    // just for good measure
    return Promise.all(data.map(fn));
});

// now to get the results:

res2.then(function(data){
    console.log(data); // [16, 32, 48, 64, 80]
});

//-------------------------------------------------------------------------------------------------------------------------------------
*/