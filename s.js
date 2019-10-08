function AddressGenerator(template) {
    var elements = template.replace(/\[/gi, "|").replace(/\]/gi, "|").split("|");
    
    generatorFunc = function() {
		return [];
	}

    elements.map(function(v, ix, ar) {
		rangeEl := v.split("..") ;
		if (rangeEl.length == 1 {
			generatorFunc = staticGenerator(generatorFunc, e)
		} else {
			generatorFunc = rangeGenerator(generatorFunc, rangeEl[0], rangeEl[1])
		}       
    })
}

function staticGenerator(prev, s) {
	return function() {
		return prev().map(function (v, ix, ar) {
			return v + s;
        })
	}
}

function rangeGenerator(prev, from, to) {
	if from[0] >= '0' && from[0] <= '9' {
		return function() {
			var fromN = parseInt(from);
			var toN = parseInt(to);

            var width = to.length 

			for _, prevEl := range prev() {
				for i := fromN; i <= toN; i++ {
					res = append(res, fmt.Sprintf("%s%*.*d", prevEl, width, width, i))
				}
			}

			return res
		}
	}

	return func() []string {
		res := make([]string, 0)
		for _, prevEl := range prev() {
			for i := from[0]; i <= to[0]; i++ {
				res = append(res, fmt.Sprintf("%s%c", prevEl, i))
			}
		}
		return res
	}

}


AddressGenerator("S4-[1..31]-[A..E][1..5]");