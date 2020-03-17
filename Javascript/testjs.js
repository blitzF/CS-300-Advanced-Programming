const fs = require('fs')
const path = require('path');


const testfolder = './astest'

/*for (let j = 0; j < process.argv.length; j++) {  
    console.log(j + ' -> ' + (process.argv[j]));
}*/
if(process.argv.length != 5)
{
	console.log('Process Exiting ! Please check your arguments !!')
	process.exit(1);
}

var modeOfFile = process.argv[2]
var filenameJS = process.argv[3]
var input = process.argv[4]
var obj = {
	enteries: []
}

if(modeOfFile == 'index')
{
	function filewalker(dir, done) {
    let results = [];

    fs.readdir(dir, function(err, list) {
        if (err) return done(err)
        var pending = list.length
        if (!pending) return done(null, results);
        list.forEach(function(file){
            file = path.resolve(dir, file)
            fs.stat(file, function(err, stat){
                if (stat && stat.isDirectory()) {
                    results.push(file)
                    filewalker(file, function(err, res){
                        results = results.concat(res);
                        if (!--pending) done(null, results);
                    })
                } else {
                	var array = fs.readFileSync(file,'utf8').toString().split("\n")
                	for(i in array) {
                	var words_in_line = array[i].split(" ")
                	for(j in words_in_line) {

                	let clean_word = words_in_line[j].replace(/[^A-Za-z0-9_]/g,"")
                	if(obj.enteries.some(function(o){return o["name"] === clean_word})) 
                	{
                		let ex = obj.enteries.find(o => o['name'] === clean_word)
                		if(ex.line_num != i && ex.location == file){
                		ex.line_num.push(i)
                		
                		}
                		if(ex.location != file){
                			obj.enteries.push({name: clean_word, location: file ,line_num: [i]})
                		}
                		
                	}
                	else {
                		obj.enteries.push({name: clean_word, location: file ,line_num: [i]})
                	}
                	var json = JSON.stringify(obj)
					fs.writeFile(filenameJS, json, (err) => {
  					if (err) {
    					console.error(err)
    					return
  					}
  					
					})
                	
                	}
				}
                	
                    results.push(file)
                    if (!--pending) done(null, results);
                }
            })
        })
    })
}

}

else if (modeOfFile == 'search')
{
	var read_dict = {}
	fs.readFile(filenameJS,'utf8',function(err,data){
		if(err) throw err
		read_dict = JSON.parse(data)
        if(read_dict.enteries.some(function(o){return o["name"] === input}))
        {
            let exx = read_dict.enteries.find(o => o['name'] === input)
            let file_loc = exx.location
            let lines_cont = exx.line_num
            var arrayy = fs.readFileSync(file_loc,'utf8').toString().split("\n")
            for(i in arrayy) {
                if(lines_cont.some(function(ele){return ele === i}))
                {
                    console.log(arrayy[i])
                }
            }


        }
        else
        {

        }

	})	
}


else {
	console.log('INVALID MODE (Either search or index)')
	process.exit(1)
}


/*filewalker(testfolder, function(err, data){
    if(err){
        throw err;
    }
   // console.log(data);
})*/