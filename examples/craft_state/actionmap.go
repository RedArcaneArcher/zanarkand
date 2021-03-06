package main

var actionMap = map[uint32]string{
	5518:   "Failed",
	252:    "Inner Quiet",
	253:    "Inner Quiet",
	254:    "Inner Quiet",
	255:    "Inner Quiet",
	256:    "Inner Quiet",
	257:    "Inner Quiet",
	258:    "Inner Quiet",
	259:    "Inner Quiet",
	260:    "Great Strides",
	261:    "Great Strides",
	262:    "Great Strides",
	263:    "Great Strides",
	264:    "Great Strides",
	265:    "Great Strides",
	266:    "Great Strides",
	267:    "Great Strides",
	4615:   "Name of the Elements",
	4616:   "Name of the Elements",
	4617:   "Name of the Elements",
	4618:   "Name of the Elements",
	4619:   "Name of the Elements",
	4620:   "Name of the Elements",
	4621:   "Name of the Elements",
	4622:   "Name of the Elements",
	4574:   "Manipulation",
	4575:   "Manipulation",
	4576:   "Manipulation",
	4577:   "Manipulation",
	4578:   "Manipulation",
	4579:   "Manipulation",
	4580:   "Manipulation",
	4581:   "Manipulation",
	4631:   "Waste Not",
	4632:   "Waste Not",
	4633:   "Waste Not",
	4634:   "Waste Not",
	4635:   "Waste Not",
	4636:   "Waste Not",
	4637:   "Waste Not",
	4638:   "Waste Not",
	4639:   "Waste Not II",
	4640:   "Waste Not II",
	4641:   "Waste Not II",
	4642:   "Waste Not II",
	4643:   "Waste Not II",
	4644:   "Waste Not II",
	19002:  "Waste Not II",
	19003:  "Waste Not II",
	19004:  "Innovation",
	19005:  "Innovation",
	19006:  "Innovation",
	19007:  "Innovation",
	19008:  "Innovation",
	19009:  "Innovation",
	19010:  "Innovation",
	19011:  "Innovation",
	19297:  "Veneration",
	19298:  "Veneration",
	19299:  "Veneration",
	19300:  "Veneration",
	19301:  "Veneration",
	19302:  "Veneration",
	19303:  "Veneration",
	19304:  "Veneration",
	19012:  "Final Appraisal",
	19013:  "Final Appraisal",
	19014:  "Final Appraisal",
	19015:  "Final Appraisal",
	19016:  "Final Appraisal",
	19017:  "Final Appraisal",
	19018:  "Final Appraisal",
	19019:  "Final Appraisal",
	100323: "Delicate Synthesis",
	100324: "Delicate Synthesis",
	100325: "Delicate Synthesis",
	100326: "Delicate Synthesis",
	100327: "Delicate Synthesis",
	100328: "Delicate Synthesis",
	100329: "Delicate Synthesis",
	100330: "Delicate Synthesis",
	100002: "Basic Touch",
	100016: "Basic Touch",
	100031: "Basic Touch",
	100076: "Basic Touch",
	100046: "Basic Touch",
	100061: "Basic Touch",
	100091: "Basic Touch",
	100106: "Basic Touch",
	100004: "Standard Touch",
	100018: "Standard Touch",
	100034: "Standard Touch",
	100078: "Standard Touch",
	100048: "Standard Touch",
	100064: "Standard Touch",
	100093: "Standard Touch",
	100109: "Standard Touch",
	100227: "Prudent Touch",
	100228: "Prudent Touch",
	100229: "Prudent Touch",
	100230: "Prudent Touch",
	100231: "Prudent Touch",
	100232: "Prudent Touch",
	100233: "Prudent Touch",
	100234: "Prudent Touch",
	100355: "Hasty Touch",
	100356: "Hasty Touch",
	100357: "Hasty Touch",
	100358: "Hasty Touch",
	100359: "Hasty Touch",
	100360: "Hasty Touch",
	100361: "Hasty Touch",
	100362: "Hasty Touch",
	100219: "Patient Touch",
	100220: "Patient Touch",
	100221: "Patient Touch",
	100222: "Patient Touch",
	100223: "Patient Touch",
	100224: "Patient Touch",
	100225: "Patient Touch",
	100226: "Patient Touch",
	100243: "Focused Touch",
	100244: "Focused Touch",
	100245: "Focused Touch",
	100246: "Focused Touch",
	100247: "Focused Touch",
	100248: "Focused Touch",
	100249: "Focused Touch",
	100250: "Focused Touch",
	100128: "Precise Touch",
	100129: "Precise Touch",
	100130: "Precise Touch",
	100131: "Precise Touch",
	100132: "Precise Touch",
	100133: "Precise Touch",
	100134: "Precise Touch",
	100135: "Precise Touch",
	100299: "Preparatory Touch",
	100300: "Preparatory Touch",
	100301: "Preparatory Touch",
	100302: "Preparatory Touch",
	100303: "Preparatory Touch",
	100304: "Preparatory Touch",
	100305: "Preparatory Touch",
	100306: "Preparatory Touch",
	100387: "Reflect",
	100388: "Reflect",
	100389: "Reflect",
	100390: "Reflect",
	100391: "Reflect",
	100392: "Reflect",
	100393: "Reflect",
	100394: "Reflect",
	100283: "Trained Eye",
	100284: "Trained Eye",
	100285: "Trained Eye",
	100286: "Trained Eye",
	100287: "Trained Eye",
	100288: "Trained Eye",
	100289: "Trained Eye",
	100290: "Trained Eye",
	100339: "Byregot's Blessing",
	100340: "Byregot's Blessing",
	100341: "Byregot's Blessing",
	100342: "Byregot's Blessing",
	100343: "Byregot's Blessing",
	100344: "Byregot's Blessing",
	100345: "Byregot's Blessing",
	100346: "Byregot's Blessing",
	100001: "Basic Synthesis",
	100015: "Basic Synthesis",
	100030: "Basic Synthesis",
	100075: "Basic Synthesis",
	100045: "Basic Synthesis",
	100060: "Basic Synthesis",
	100090: "Basic Synthesis",
	100105: "Basic Synthesis",
	100203: "Careful Synthesis",
	100204: "Careful Synthesis",
	100205: "Careful Synthesis",
	100206: "Careful Synthesis",
	100207: "Careful Synthesis",
	100208: "Careful Synthesis",
	100209: "Careful Synthesis",
	100210: "Careful Synthesis",
	100363: "Rapid Synthesis",
	100364: "Rapid Synthesis",
	100365: "Rapid Synthesis",
	100366: "Rapid Synthesis",
	100367: "Rapid Synthesis",
	100368: "Rapid Synthesis",
	100369: "Rapid Synthesis",
	100370: "Rapid Synthesis",
	100235: "Focused Synthesis",
	100236: "Focused Synthesis",
	100237: "Focused Synthesis",
	100238: "Focused Synthesis",
	100239: "Focused Synthesis",
	100240: "Focused Synthesis",
	100241: "Focused Synthesis",
	100242: "Focused Synthesis",
	100315: "Intensive Synthesis",
	100316: "Intensive Synthesis",
	100317: "Intensive Synthesis",
	100318: "Intensive Synthesis",
	100319: "Intensive Synthesis",
	100320: "Intensive Synthesis",
	100321: "Intensive Synthesis",
	100322: "Intensive Synthesis",
	100403: "Groundwork",
	100404: "Groundwork",
	100405: "Groundwork",
	100406: "Groundwork",
	100407: "Groundwork",
	100408: "Groundwork",
	100409: "Groundwork",
	100410: "Groundwork",
	100379: "Muscle Memory",
	100380: "Muscle Memory",
	100381: "Muscle Memory",
	100382: "Muscle Memory",
	100383: "Muscle Memory",
	100384: "Muscle Memory",
	100385: "Muscle Memory",
	100386: "Muscle Memory",
	100331: "Brand of the Elements",
	100332: "Brand of the Elements",
	100333: "Brand of the Elements",
	100334: "Brand of the Elements",
	100335: "Brand of the Elements",
	100336: "Brand of the Elements",
	100337: "Brand of the Elements",
	100338: "Brand of the Elements",
	100010: "Observe",
	100023: "Observe",
	100040: "Observe",
	100053: "Observe",
	100070: "Observe",
	100082: "Observe",
	100099: "Observe",
	100113: "Observe",
	100003: "Masters Mend",
	100017: "Masters Mend",
	100032: "Masters Mend",
	100047: "Masters Mend",
	100062: "Masters Mend",
	100077: "Masters Mend",
	100092: "Masters Mend",
	100107: "Masters Mend",
	100371: "Tricks of the Trade",
	100372: "Tricks of the Trade",
	100373: "Tricks of the Trade",
	100374: "Tricks of the Trade",
	100375: "Tricks of the Trade",
	100376: "Tricks of the Trade",
	100377: "Tricks of the Trade",
	100378: "Tricks of the Trade",
}
