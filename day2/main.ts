const text = await Deno.readTextFile('./input.txt');
const lines = text.split('\n');

const safeLines = lines.filter(line => {
	const elements = line.split(' ').map(e => parseInt(e));

	if (elements.length < 2) {
		return true;
	}

	const shouldBeIncreasing = elements[1] > elements[0];
	for (let i = 1; i < elements.length; i++) {
		const diff = Math.abs(elements[i] - elements[i - 1]);

		if (diff < 1 || diff > 3) {
			return false;
		}

		if (shouldBeIncreasing && elements[i] <= elements[i - 1]) {
			return false;
		}
		if (!shouldBeIncreasing && elements[i] >= elements[i - 1]) {
			return false;
		}
	}

	return true;
});

console.log('There are', safeLines.length, 'safe lines');
