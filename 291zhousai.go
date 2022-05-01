package main

import (
	"fmt"
	"math"
)

func removeDigit(number string, digit byte) string {
	var maxString string
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			var ss string
			if i == len(number)-1 {
				ss = number[:i]
			} else {
				ss = number[:i] + number[i+1:]
			}
			if maxString < ss {
				maxString = ss
			}
		}
	}
	return maxString
}

func minimumCardPickup(cards []int) int {
	minL := math.MaxInt32
	flag := false
	m := make(map[int]int, 1e5)
	for i := 0; i < len(cards); i++ {
		if _, ok := m[cards[i]]; ok {
			flag = true
			if minL > i-m[cards[i]]+1 {
				minL = i - m[cards[i]] + 1
			}
		}
		m[cards[i]] = i
	}
	if !flag {
		minL = -1
	}
	return minL
}

func main() {
	//fmt.Println(removeDigit("7795478535679443616467964135298543163376223791274561861738666981419251859535331546947347395531332878", '5'))
	//fmt.Println(removeDigit("1231", '1'))
	//fmt.Println(removeDigit("551", '5'))
	fmt.Println(minimumCardPickup([]int{4901, 9272, 2244, 1494, 1455, 3761, 5039, 3434, 5983, 578, 6022,
		8219, 2561, 8548, 733, 6601, 9344, 5654, 9043, 1162, 7252, 7424, 9467, 8838, 8618, 2076, 2538, 7581,
		2199, 9472, 8911, 3144, 4465, 3170, 5740, 4176, 6738, 7723, 672, 8888, 4057, 184, 2527, 5822, 2505,
		3988, 7975, 4356, 3531, 830, 7282, 2372, 5248, 6336, 6015, 4468, 5648, 9591, 3520, 8296, 4073, 5909,
		99, 3312, 8006, 8866, 8582, 7991, 3110, 2113, 2956, 5690, 2475, 3091, 1459, 751, 659, 1529, 6332, 3953,
		6225, 1023, 88, 6025, 3275, 9317, 6967, 262, 7622, 4259, 7618, 7608, 3276, 9574, 5608, 2426, 8762, 7407,
		5443, 9860, 7245, 1308, 2897, 7253, 1377, 2838, 5789, 2856, 4881, 593, 6397, 4138, 2304, 6342, 6375, 8014,
		6735, 6189, 880, 6002, 7056, 8097, 9094, 8094, 7652, 4474, 3801, 9023, 9409, 2354, 4588, 7331, 8348, 1385,
		1291, 8593, 8127, 1047, 5113, 763, 8363, 8179, 7444, 8627, 5460, 7913, 5612, 1629, 7233, 5240, 238, 1101,
		5249, 4242, 5816, 9296, 2417, 1520, 4650, 8138, 5750, 4871, 2013, 4091, 4821, 8381, 9054, 2836, 5663, 9477, 9272, 4799, 8557, 4167, 1138, 4107, 7942, 5866, 1466, 5447, 3617, 1016, 5379, 1677, 9575, 4641, 7699, 9362, 2493, 3693, 4058, 7290, 7835, 4304, 265, 2220, 3416, 3698, 4281, 7460, 1991, 2228, 4631, 8181, 1344, 5306, 6343, 8577, 620, 3833, 4494, 2075, 7304, 728, 3310, 898, 9899, 9870, 1684, 4546, 9979, 9272, 7953, 2010, 9144, 8068, 1880, 3834, 2536, 5194, 9137, 2718, 1251, 3477, 1743, 5766, 7713, 5903, 9492, 6291, 6821, 6999, 3833, 3420, 5891, 9164, 595, 812, 8929, 5221, 6599, 5956, 9882, 9219, 9549, 6940, 2869, 476, 7217, 997, 561, 7402, 7469, 2978, 4049, 9521, 7796, 3348, 3577, 9131, 9923, 1996, 7794, 9242, 8238, 1789, 7621, 7434, 6713, 658, 3946, 8863, 5808, 9534, 9365, 9670, 807, 3909, 4915, 3180, 4668, 1346, 2986, 5756, 2844, 4182, 1470, 5380, 9653, 1057, 5155, 9986, 8903, 8577, 790, 2600, 1448, 7867, 4942, 7023, 1721, 2770, 2521, 9423, 2104, 2441, 5255, 9704, 6958, 213, 9443, 3359, 7654, 608, 7041, 9391, 4340, 6209, 3258, 3634, 6617, 1080, 1268, 5710, 8562, 6607, 9267, 8840, 3387, 8279, 9961, 3406, 5466, 6469, 95, 1707, 3198, 8168, 1827, 5993, 9942, 5614, 6930, 7698, 8409, 2277, 6920, 8733, 2463, 4418, 6097, 156, 5564, 4661, 2635, 6296, 1283, 3817, 967, 2554, 7188, 9721, 1984, 4326, 5340, 4453, 1273, 915, 4923, 3416, 7925, 8715, 7924, 4699, 7105, 6393, 900, 5517, 712, 3121, 2068, 3954, 6012, 7158, 5339, 2331, 2234, 2045, 8489, 1310, 7827, 507, 1421, 6215, 405, 6856, 4787, 212, 783, 5622, 747, 558, 4297, 4372, 8853, 2138, 2183, 5081, 1687, 224, 4599, 1712, 3510, 1854, 3926, 4433, 6791, 4879, 2067, 5307, 5488, 2685, 7159, 7086, 1072, 2671, 3312, 8594, 2229, 8348, 3114, 3312, 379, 7035, 9404, 2691, 2295, 8143, 305, 7747, 8091, 9256, 5686, 9292, 1848, 2790, 3243, 4834, 8889, 3743, 3952, 571, 6429, 5077, 6277, 5126, 8225, 8354, 2047, 1118, 6446, 2214, 7425, 9520, 8843, 7157, 5922, 115, 3419, 1790, 8924, 2486, 1587, 1608, 3818, 2409, 7794, 6274, 1933, 2345, 9921, 1790, 8912, 681, 9884, 1443, 3850, 4388, 2323, 4247, 3444, 656, 472, 7554, 8644, 2170, 613, 3241, 6332, 2786, 7069, 1478, 4353, 7458, 7614, 1809, 5768, 6613, 4158, 9199, 221, 5841, 2134, 5108, 3777, 7854, 2604, 7317, 622, 5653, 7341, 4256, 4358, 9921, 2263, 5268, 8347, 8462, 3105, 6895, 7085, 8391, 700, 9416, 8788, 4600, 4472, 3961, 9975, 7358, 4311, 6322, 4996, 3001, 6695, 1290, 6182, 340, 1888, 2275, 1582, 3286, 1840, 6458, 7757, 9546, 8825, 3776, 2785, 2901, 440, 9494, 2302, 676, 5870, 6487, 6685, 6372, 9851, 3472, 2501, 1268, 5388, 8058, 4021, 2343, 5546, 8874, 8330, 498, 8285, 5680, 1840, 2357, 7252, 9587, 196, 2692, 3197, 4757, 4967, 2873, 8955, 269, 7540, 8178, 8924, 9690, 3492, 6848, 7242, 4018, 3237, 2391, 6145, 9859, 4361, 8163, 2310, 7699, 346, 1788, 123, 8810, 8305, 9670, 2997, 4635, 4772, 3113, 5413, 2202, 4716, 9118, 4713, 5954, 3254, 9915, 4413, 7655, 8523, 6065, 5216, 2936, 2286, 459, 7328, 9102, 7399, 5229, 4833, 6907, 8521, 1904, 8186, 3081, 1157, 3847, 2668, 1093, 12, 1826, 2510, 6148, 1583, 3185, 5900, 4000, 7244, 293, 8090, 1831, 3963, 6785, 8638, 8550, 2578, 1652, 3853, 3877, 4732, 320, 7970, 3008, 1361, 4056, 5885, 9241, 3594, 3555, 1704, 6357, 4418, 3429, 4242, 1468, 9634, 5443, 9073, 8653, 6893, 1251, 9655, 8654, 9372, 7904, 2420, 5995, 5987, 3836, 224, 4721, 9512, 3180, 3371, 1688, 2418, 727, 274, 9291, 4774, 4669, 9759, 9033, 8739, 2997, 848, 2469, 1355, 6978, 2458, 5188, 2881, 9113, 2339, 549, 9990, 3398, 9482, 188, 3235, 2849, 3145, 8208, 7929, 5681, 6520, 2300, 3438, 9296, 7247, 788, 9596, 4152, 4432, 9859, 6096, 9158, 883, 9617, 3770, 3784, 9454, 8984, 9511, 7602, 6274, 6035, 1917, 3192, 3670, 1944, 7587, 7583, 2043, 506, 145, 7068, 142, 4011, 5464, 2581, 2950, 6424, 6948, 3784, 7990, 6425, 6913, 3866, 2708, 6245, 5344, 3894, 5398, 2728, 8687, 966, 5813, 769, 5102, 9353, 1819, 5149, 6349, 881, 1369, 3666, 2565, 5139, 3476, 299, 2956, 8878, 9189, 6067, 8691, 9931, 1049, 1481, 5491, 8031, 6033, 872, 1487, 3558, 664, 8607, 5562, 6002, 151, 9258, 5521, 5698, 3875, 243, 1124, 8364, 5452, 7220, 8433, 5421, 5058, 3194, 9215, 2220, 9273, 9745, 1283, 5347, 8074, 9980, 3144, 733, 1400, 7101, 6176, 1585, 5663, 7434, 3935, 7667, 8592, 8551, 1275, 5709, 9421, 6354, 5988, 7358, 8581, 5537, 5322, 2318, 6755, 6972, 1413, 8546, 1188, 4165, 6629, 1987, 5022, 8467, 7882, 3313, 1581, 4221, 5220, 5273, 4679, 4846, 1120, 4500, 2602, 9952, 8897, 5233, 2721, 5284, 328, 911, 7427, 2608, 4413, 5313, 4826, 6435, 1706, 6217, 377, 9474, 4171, 658, 7115, 3361, 8888, 1664, 8961, 3972, 2765, 5744, 1923, 3528, 4653, 1375, 6774, 5296, 5911, 2854, 4135, 3529, 7198, 2810, 5718, 4655, 2566, 1232, 9702, 3454, 5832, 9205, 8094, 3285, 8624, 8959, 125, 2559, 568, 2708, 2154, 9612, 7111, 8266, 7994, 2376, 343, 5154, 8348, 9599, 6729, 304, 8433, 8391, 4716, 7535, 5504, 115, 4739, 6459, 567, 2175, 3690, 6009, 5175, 3720, 342, 6105, 4911, 2948, 7575, 8808, 3799, 783, 6366, 2631, 8063, 4716, 3449, 8541, 7189, 2646, 2156, 7721, 7450, 4616, 251, 5824, 9090, 6342, 6750, 2647, 374, 5383, 2815, 3276, 5037, 6477, 2225, 6873, 17, 9672, 9441, 8470, 9198, 1632, 8191, 9157, 834, 2733, 1358, 5492, 7786, 9396, 7729, 8915, 8489, 179, 4216, 3775, 508, 2799, 6588, 8959, 3521, 4890, 5201, 4512, 3630, 282, 6445, 1237, 1282, 6873, 3021, 5405, 5847, 5133, 5524, 9846, 1577, 5455, 4028, 9518, 9419, 6723, 8444, 986, 5275, 3829, 2694, 2114, 6534, 1777, 9255, 5625, 6892, 6500, 4836, 5805, 5130, 2803, 6212, 2110, 327, 5974, 5461, 4712, 4231, 847, 8039, 2048, 208, 7094, 3901, 9942, 4496, 5196, 4436, 6293, 1160, 3344, 513, 7171, 1538, 104, 4877, 7643, 429, 5069, 7929, 1994, 3200, 7483, 2492, 7028, 3232, 646, 9172, 5668, 4531, 3757, 3824, 194, 5271, 9507, 6821, 2206, 9451, 3148, 9366, 744, 7171, 1324, 2707, 2610, 7227, 4359, 3309, 586, 2047, 8089, 3592, 807, 9519, 3774, 1055, 9872, 9295, 4026, 9206, 3083, 4976, 6501, 3564, 9460, 3536, 9266, 7212, 2350, 2363, 4117, 4741, 7282, 689, 8769, 2083, 7702, 9193, 5864, 8223, 9359, 8654, 4070, 3700, 1628, 1479, 7573, 7740, 6922, 7985, 9349, 2383, 5083, 9686, 5429, 5688, 6523, 1830, 4257, 5611, 6997, 9009, 3774, 4331, 3250, 4145, 3616, 3775, 4161, 3988, 1329, 6962, 6942, 5521, 2240, 9899, 8013, 7102, 9114, 7842, 5456, 4533, 7103, 5627, 2724, 2712, 5322, 5628, 4925, 7713, 8464, 5792, 4146, 7599, 9488, 4286, 9655, 2724, 9940, 1215, 2595, 7641, 6625, 7790, 7364, 4717, 4005, 634, 444, 1762, 4098, 7298, 7091, 2071, 8779, 3292, 9174, 7020, 6398, 6484, 6049, 7641, 4334, 5235, 4092, 551, 8452, 1525, 1991, 2039, 9730, 463, 7492, 802, 1521, 7879, 1569, 8485, 7893, 3397, 237, 7394, 2350, 5890, 3667, 6475, 7954, 5087, 1684, 1947, 2935, 8931, 4466, 2026, 9602, 9019, 42, 8494, 5909, 6463, 4593, 1132, 7346, 7788, 6753, 1851, 6823, 4105, 6332, 9826, 7879, 2534, 481, 1958, 1299, 1403, 4179, 6978, 2324, 1821, 5155, 9866, 1093, 4375, 8077, 6852, 6373, 3774, 8741, 8475, 2148, 1351, 5730, 159, 3559, 5468, 8616, 6861, 6779, 770, 5769, 9758, 4108, 7674, 7347, 346, 7772, 5085, 4505, 6785, 9275, 3763, 9812, 1052, 6476, 2710, 8073, 8033, 8816, 5550, 6391, 4053, 3802, 7430, 6263, 3038, 9477, 3033, 9085, 7565, 7018, 1244, 9574, 6272, 989, 4401, 9454, 3710, 1239, 2424, 9869, 7316, 9732, 3959, 4959, 9020, 3551, 7927, 5920, 3481, 8216, 2065, 9070, 3454, 5708, 1403, 8546, 2848, 5317, 8631, 9314, 7023, 9745, 2227, 1779, 9201, 5790, 6196, 3219, 7582, 4509, 3327, 4178, 4867, 6875, 4742, 4277, 5261, 9013, 6939, 8381, 2249, 6877, 2043, 4226, 3930, 8348, 3265, 5991, 4375, 4333, 5792, 3822, 2723, 7878, 1192, 8279, 8622, 3417, 1256, 2938, 8415, 3865, 4628, 2956, 4705, 1217, 149, 4764, 7923, 4622, 2958, 2713, 4207, 2143, 6291, 2879, 3537, 828, 7471, 2701, 6950, 682, 7020, 5074, 6344, 968, 6381, 359, 1610, 3307, 4484, 4296, 3687, 3344, 2958, 7002, 5606, 5184, 4773, 3484, 7648, 560, 9564, 584, 4441, 9430, 7802, 1440, 6872, 473, 4437, 3768, 595, 2034, 6380, 9583, 7975, 7684, 838, 1133, 4637, 3675, 7783, 5658, 1868, 4854, 9646, 2257, 9927, 5489, 8471, 5197, 9964, 3115, 3942, 301, 2196, 3854, 9873, 2098, 7187, 6315, 2690, 2426, 8164, 904, 6544, 9048, 8653, 6405, 2963, 4486, 2389, 6481, 1445, 6305, 883, 8725, 8535, 6420, 6200, 4271, 2487, 3054, 1597, 1515, 6671, 4579, 3726, 308, 8830, 1589, 3103, 8231, 100, 479, 2164, 6290, 2919, 9784, 2490, 3193, 7585, 3428, 2443, 9054, 4349, 9687, 4962, 6348, 5218, 8594, 1327, 648, 5695, 3813, 1803, 6448, 943, 1129, 250, 2509, 2950, 9880, 8636, 534, 9418, 2125, 7907, 7348, 2691, 8819, 2850, 2634, 7664, 883, 4708, 4425, 8659, 7217, 6482, 4562, 6798, 8696, 8741, 205, 3745, 62, 5190, 2959, 5922, 9962, 3442, 6899, 8326, 3856, 4431, 6873, 9575, 504, 6419, 4332, 5906, 2591, 6685, 7989, 2331, 1077, 46, 6219, 858, 616, 1467, 2559, 4887, 1197, 6655, 3950, 6155, 8557, 5828, 5165, 8931, 5608, 5243, 5150, 6168, 4007, 3902, 6940, 3845, 8926, 345, 8221, 1596, 5550, 6047, 1359, 2707, 3914, 9084, 9151, 7223, 460, 5684, 8424, 9160, 1075, 4252, 8244, 5347, 1818, 8692, 9805, 7551, 5120, 4700, 6172, 3158, 4137, 6044, 727, 3979, 971, 1842, 5285, 9226, 963, 6644, 50, 7631, 5335, 8308, 1815, 8770, 9100, 6694, 7558, 9088, 9523, 3777, 9785, 1770, 4783, 9685, 5001, 2155, 3528, 3754, 6472, 4474, 2044, 6950, 162, 6468, 3866, 7281, 8736, 410, 3149, 7254, 6772, 1212, 6734, 8789, 9585, 2748, 4826, 4997, 9073, 4573, 8120, 8929, 3736, 7157, 3828, 4490, 5420, 61, 2161, 5862, 4476, 2700, 7925, 1395, 9743, 2492, 4902, 1937, 4702, 1241, 5966, 5747, 8404, 7505, 2387, 6662, 5544, 9306, 2302, 851, 9434, 1, 4537, 5418, 3937, 8650, 30, 1069, 1131, 4682, 6726, 4399, 2930, 9048, 5678, 1018, 480, 5469, 1998, 1796, 7976, 6483, 5959, 4086, 7816, 4725, 6937, 9431, 2279, 9101, 4494, 4773, 5900, 4324, 2389, 1649, 3866, 7733, 3715, 4073, 145, 9876, 5190, 8587, 6378, 3932, 6860, 2579, 7832, 6601, 5997, 5902, 7462, 6896, 4772, 885, 1538, 767, 1845, 285, 9985, 7792, 9585, 2311, 9449, 7097, 8951, 7581, 6275, 8409, 820, 3547, 1695, 9068, 2130, 7238, 4104, 4891, 5999, 7954, 4319, 1667, 5340, 6432, 5562, 4040, 320, 1661, 1206, 5805, 5017, 9197, 6782, 2917, 8272, 7310, 8416, 2887, 1754, 3083, 6875, 9335, 4749, 1589, 4026, 9566, 7790, 7267, 2296, 2728, 1532, 2571, 6273, 2539, 1675, 3307, 6796, 862, 4668, 5216, 1923, 4037, 2313, 8482, 1009, 4285, 8077, 1213, 2398, 4788, 7353, 143, 9225, 4489, 6488, 6972, 3708, 4124, 7845, 4021, 613, 1839, 1121, 6155, 5131, 9480, 5281, 6128, 7091, 335, 7977, 5680, 7697, 9946, 2425, 7405, 7322, 6592, 5151, 5187, 6390, 4961, 4262, 4263, 8834, 6559, 110, 3229, 3926, 1081, 7527, 2786, 6869, 9798, 6165, 5180, 1869, 3681, 1363, 6137, 7098, 768, 4549, 3336, 6572, 919, 2600, 8893, 3581, 2784, 2822, 8570, 8847, 6946, 4629, 359, 4646, 1691, 7164, 1859, 4445, 2853, 1114, 2125, 2897, 1098, 7397, 5492, 3055, 6620, 841, 2410, 7270, 9228, 4225, 3528, 3260, 1553, 7547, 660, 7112, 5511, 5836, 6249, 3549, 4165, 3732, 9947, 4608, 6694, 8588, 8962, 3869, 6076, 1500, 1337, 7955, 1591, 4957, 8771, 8617, 3562, 8875, 2489, 8314, 3668, 2276, 1809, 4328, 1864, 2113, 7281, 7498, 3794, 4616, 8724, 6932, 1760, 7957, 5781, 6301, 7828, 2379, 2215, 6964, 3083, 9509, 3935, 3420, 4014, 9152, 3595, 7894, 6207, 7062, 352, 4504, 9333, 8651, 9698, 8093, 3181, 2249, 2227, 1749, 6773, 2368, 5850, 416, 2876, 8035, 8730, 5550, 9850, 6837, 4951, 3421, 1468, 8857, 4152, 2727, 4407, 2743, 2302, 8199, 6490, 761, 5265, 5343, 8815, 7227, 9900, 2403, 9554, 402, 8684, 400, 7889, 4570, 8012, 2801, 827, 4695, 6559, 5360, 4071, 1388, 719, 9152, 1513, 2381, 2734, 4833, 9760, 6942, 8060, 6265, 5169, 8955, 5193, 6053, 1995, 7669, 3812, 3194, 4319, 4682, 3888, 8653, 5611, 1275, 521, 2818, 1169, 9701, 6703, 402, 7753, 3922, 997, 2612, 7424, 7939, 8914, 8506, 1800, 9535, 8762, 3401, 9614, 8786, 9205, 4770, 3730, 3154, 1995, 151, 6062, 6924, 252, 4650, 4720, 8598, 3483, 4410, 1128, 9336, 2343, 242, 778, 2776, 1436, 7239, 6067, 6830, 2128, 1735, 6396, 1087, 3649, 5924, 3319, 7531, 188, 4594, 5596, 107, 9492, 326, 7979, 9433, 5056, 8940, 5800, 8260, 9708, 6218, 4430, 6766, 7234, 5870, 2930, 6065, 7057, 1305, 8707, 8094, 8551, 3163, 8538, 8717, 6810, 1176, 7254, 2996, 3913, 636, 9660, 8354, 5755, 2859, 1495, 1096, 1546, 9719, 6310, 6371, 7786, 5656, 8660, 7768, 7407, 9173, 2025, 416, 6648, 843, 4039, 8250, 5622, 1419, 4849, 8422, 7116, 8117, 221, 6296, 372, 8771, 4540, 665, 2610, 3507, 7825, 4019, 6555, 6014, 6095, 5838, 5780, 2934, 8429, 565, 9691, 7163, 58, 2297, 4976, 547, 7629, 886, 1253, 1810, 9644, 7403, 1670, 7406, 2825, 9351, 4245, 8168, 2087, 6892, 2447, 7954, 9192, 2159, 9084, 4612, 883, 2920, 3493, 950, 8114, 4560, 791, 7055, 1889, 7853, 911, 1960, 3416, 4528, 1, 3491, 2312, 5789, 3757, 6046, 5079, 6016, 1818, 6885, 6902, 8479, 1436, 6201, 670, 4066, 245, 8751, 4764, 6262, 7040, 1774, 5042, 5641, 3522, 3304, 4202, 4059, 270, 970, 6382, 6122, 7150, 1362, 3332, 8554, 3442, 1833, 914, 684, 4647, 1185, 3835, 4364, 8655, 515, 2562, 6633, 9807, 3050, 8386, 7406, 7835, 234, 1021, 9804, 3410, 1857, 900, 902, 1490, 8572, 4093, 3037, 3956, 9326, 4178, 9914, 1470, 1997, 3838, 6358, 6496, 3999, 7484, 1558, 4505, 1096, 7508, 5904, 2525, 4269, 414, 5844, 9637, 8231, 6753, 6977, 6748, 5067, 7866, 5170, 3340, 8171, 2714, 4512, 9714, 1738, 7629, 4949, 4168, 636, 2150, 284, 5512, 6166, 6711, 3067, 2729, 4451, 290, 5258, 678, 5485, 7764, 8014, 4001, 7672, 7647, 7338, 3264, 8587, 1254, 6800, 3316, 4434, 1145, 6221, 3954, 5786, 4807, 1361, 7049, 5519, 1934, 2626, 8699, 4069, 1384, 9200, 595, 649, 6570, 9780, 8859, 5929, 1567, 6480, 9317, 3731, 2718, 8788, 7504, 8292, 9255, 9573, 9837, 8867, 3812, 9745, 3454, 5264, 278, 4148, 3434, 7690, 1666, 8239, 213, 496, 7543, 5635, 8446, 7256, 6015, 6782, 8175, 8562, 3607, 8885, 3420, 5283, 6613, 2985, 789, 1164, 4407, 3126, 3460, 2693, 4550, 5664, 6981, 8787, 9027, 3052, 1989, 7832, 1772, 7842, 2441, 9965, 7938, 4439, 3066, 3888, 106, 4349, 905, 5378, 9893, 2996, 4815, 946, 5164, 7347, 5185, 8735, 3965, 9285, 5204, 5080, 497, 2825, 2436, 1984, 2091, 4001, 1577, 4462, 1520, 5044, 2871, 4592, 1075, 1511, 5884, 5818, 9289, 539, 3094, 7436, 4973, 3172, 1380, 5037, 8210, 6402, 9776, 7238, 3032, 3735, 7545, 4240, 4596, 1440, 2726, 8714, 7857, 5528, 7698, 66, 6082, 933, 3323, 1646, 5579, 5640, 9938, 8902, 4623, 8067, 1444, 8674, 3848, 6952, 7597, 4203, 8381, 3322, 1588, 5117, 8964, 1305, 689, 4436, 785, 7440, 8610, 8478, 5909, 6653, 2801, 4100, 9829, 1402, 8759, 1237, 9881, 9060, 118, 9204, 8821, 6303, 6822, 7059, 2325, 354, 4459, 392, 7095, 1963, 8900, 4704, 7971, 3549, 6075, 6634, 62, 6841, 3246, 4584, 7770, 2599, 4346, 1220, 5786, 1474, 1416, 4566, 319, 3476, 2207, 1745, 2603, 5023, 8231, 7939, 6682, 6531, 4431, 6566, 7655, 9125, 9127, 8221, 6911, 8073, 718, 8038, 1724, 7893, 7252, 8446, 8366, 8781, 6700, 1194, 7758, 5984, 5406, 1763, 3243, 6364, 1087, 6890, 8617, 7972, 3608, 5563, 7857, 5136, 6876, 3799, 2728, 5729, 6056, 84, 1990, 3167, 5273, 4367, 2372, 6548, 6835, 7156, 1528, 7122, 1260, 6768, 957, 5612, 539, 2443, 9499, 3256, 4221, 3763, 6660, 5314, 5321, 2361, 2593, 6603, 8012, 1295, 124, 3145, 6146, 4707, 9894, 3784, 8643, 5942, 5870, 9991, 1497, 8762, 5407, 5081, 959, 4277, 4926, 8668, 7644, 6427, 9188, 310, 337, 9881, 1386, 2366, 1804, 9799, 6026, 937, 2254, 834, 8651, 3523, 6790, 2382, 8790, 3682, 360, 8881, 3341, 2728, 8371, 1119, 4488, 3186, 5625, 7577, 6376, 2324, 8318, 4816, 7468, 2099, 6043, 4568, 702, 5269, 9397, 8343, 1439, 4970, 8445, 9539, 8531, 9682, 5125, 6112, 4191, 2209, 2687, 7643, 6142, 6280, 8777, 3871, 2030, 3682, 239, 1022, 4719, 7526, 2896, 975, 4343, 2504, 6919, 8092, 8943, 7836, 5960, 5189, 9758, 5256, 5219, 1680, 8183, 2849, 3225, 2536, 3991, 7955, 8594, 2977, 882, 5898, 5580, 1143, 5351, 5959, 6938, 4151, 2735, 9670, 231, 6028, 7086, 9024, 9197, 9210, 1022, 4593, 5968, 66, 676, 2848, 9568, 9046, 7389, 6368, 6060, 5902, 4164, 5738, 5976, 837, 4451, 584, 3644, 2344, 7783, 8744, 5007, 9519, 7037, 5761, 1788, 5181, 7060, 8617, 2529, 4779, 7280, 2553, 7774, 3679, 2950, 9506, 1245, 3631, 3412, 8610, 1265, 7011, 4165, 1436, 297, 421, 6683, 369, 9095, 1506, 4783, 6237, 6295, 7778, 3985, 6534, 2031, 415, 1531, 1019, 3491, 5814, 8746, 9889, 8717, 8783, 9274, 7494, 6398, 4975, 8010, 4072, 6331, 6535, 8164, 4918, 4606, 1108, 1296, 1244, 7854, 6774, 9835, 8569, 7064, 3512, 4506, 412, 9595, 4421, 6943, 8033, 8252, 9411, 4756, 6711, 6760, 897, 5127, 7253, 8024, 1018, 8003, 3818, 4833, 4364, 5656, 9984, 3024, 5882, 4256, 7857, 3375, 3430, 2007, 7405, 7553, 1800, 9814, 2788, 1936, 4205, 5751, 5833, 2532, 1264, 2945, 1678, 4010, 9293, 8776, 9139, 882, 533, 9327, 9311, 6346, 844, 1072, 8540, 7573, 3820, 8962, 4598, 4360, 4869, 1288, 7226, 1508, 6681, 9079, 2285, 9534, 4559, 8629, 4341, 3175, 6824, 7606, 4000, 1353, 7836, 5404, 3865, 6371, 6492, 6002, 6518, 8926, 9259, 8845, 3671, 5497, 3918, 9471, 610, 4068, 3794, 175, 517, 3246, 138, 50, 767, 1051, 7725, 4646, 6122, 7763, 2672, 459, 7641, 8656, 9813, 4715, 5091, 5572, 501, 2784, 7702, 1463, 6057, 3066, 8575, 6658, 9720, 5454, 2361, 1272, 923, 2084, 1100, 2088, 1361, 1337, 6218, 9904, 9951, 7285, 9693, 1127, 1386, 6325, 8632, 5614, 6417, 1529, 3247, 9101, 2729, 8445, 4403, 7490, 4645, 960, 335, 6604, 5976, 2835, 3433, 287, 4388, 6528, 703, 8899, 5665, 1252, 4546, 7854, 3208, 1996, 4403, 3412, 3917, 6070, 4886, 7704, 2797, 7475, 6891, 6801, 6812, 4892, 5217, 9083, 6258, 4791, 4680, 2243, 9056, 6373, 4258, 1971, 8787, 60, 4970, 2982, 286, 1956, 9981, 6422, 7210, 8490, 4376, 7886, 8260, 7693, 163, 6039, 2530, 9097, 8310, 5487, 9936, 2595, 1926, 9008, 11, 5914, 8542, 9131, 7698, 8152, 5454, 95, 8060, 7587, 1543, 4586, 9238, 7810, 1801, 5410, 3343, 9317, 5311, 4194, 5815, 3104, 709, 6309, 5501, 3433, 895, 806, 9570, 4219, 2904, 7228, 8747, 5621, 3552, 7326, 8278, 3615, 8493, 4809, 9676, 2214, 843, 3607, 5254, 9914, 9883, 7437, 5813, 2303, 3065, 7122, 9673, 9709, 8381, 5687, 7703, 3606, 9725, 1306, 3688, 8631, 5966, 5197, 5443, 1710, 7019, 9903, 4607, 3943, 1603, 5434, 4850, 1485, 7240, 7059, 7126, 7918, 1018, 111, 8207, 3072, 5698, 9122, 5897, 797, 561, 6624, 3894, 361, 9029, 3580, 8587, 2559, 8822, 2639, 1520, 5138, 1869, 5274, 8794, 795, 1895, 2239, 8311, 2835, 6853, 8262, 2646, 9815, 4705, 9920, 5861, 2786, 274, 9480, 9549, 1004, 3903, 8032, 6814, 9099, 2270, 2989, 6621, 5535, 5358, 3353, 7569, 8156, 7493, 4400, 1176, 8405, 4805, 1584, 9284, 1037, 3867, 34, 7411, 2402, 8116, 2248, 2981, 8609, 8781, 1297, 5726, 551, 5156, 5904, 5496, 1513, 974, 3517, 2754, 465, 2083, 7259, 9502, 7848, 9058, 4893, 7673, 4666, 5029, 4690, 8347, 8984, 383, 8038, 1994, 4548, 9350, 375, 17, 866, 5162, 8764, 7327, 3770, 6263, 3425, 960, 8865, 2181, 4560, 9285, 1107, 7942, 962, 8004, 3153, 4129, 9276, 2647, 2661, 8005, 1889, 9184, 6004, 8087, 866, 322, 2028, 5755, 7614, 5012, 4567, 1277, 4925, 4014, 3251, 7121, 3718, 7542, 9258, 1006, 6633, 1824, 4092, 7229, 80, 2152, 8765, 9994, 4841, 3613, 5478, 822, 5933, 946, 718, 6228, 6405, 4480, 7301, 2831, 7606, 894, 3259, 4565, 3072, 5951, 2897, 4712, 4961, 5900, 6571, 2218, 3868, 2781, 3058, 9069, 7000, 7559, 2511, 8327, 2486, 7197, 4306, 7708, 4438, 3302, 4466, 7230, 4978, 4046, 1499, 1832, 4339, 470, 4024, 9710, 3211, 4555, 764, 7282, 4594, 3993, 8290, 4644, 7565, 8977, 174, 4817, 2896, 8605, 4520, 1997, 3813, 1965, 9645, 9217, 9995, 2658, 9513, 6463, 9579, 958, 650, 2490, 233, 7083, 4154, 1599, 9732, 8012, 2052, 7815, 2310, 433, 3595, 9282, 147, 2077, 5792, 1285, 8552, 2349, 7816, 697, 9456, 4305, 4580, 7656, 722, 3618, 2773, 6905, 7936, 3009, 1867, 1464, 3356, 8542, 876, 7062, 1642, 6304, 6001, 5251, 5431, 4759, 4507, 6004, 3370, 4156, 4728, 7559, 5979, 8826, 9503, 8550, 7410, 5335, 8799, 4088, 8998, 9168, 4987, 163, 2504, 220, 191, 6500, 6652, 6672, 3341, 1079, 5449, 2870, 1979, 1090, 1829, 3132, 3879, 9593, 1554, 2105, 5398, 751, 761, 436, 25, 773, 6163, 1368, 9780, 5249, 9078, 9465, 6814, 6011, 9436, 5201, 7214, 5955, 5073, 4018, 6272, 1774, 4216, 2011, 642, 3009, 2337, 949, 9338, 5125, 3033, 2428, 3176, 4903, 9445, 2210, 1378, 5163, 3166, 3364, 7079, 3546, 6155, 8659, 7488, 2257, 1948, 2600, 7275, 5469, 874, 698, 5976, 6865, 7775, 3617, 3577, 5229, 3369, 5048, 9457, 8873, 2900, 4804, 479, 8888, 4032, 6015, 1223, 9660, 5899, 1935, 6188, 3425, 1672, 6289, 8969, 407, 7871, 8678, 8066, 5718, 8199, 9350, 6023, 8172, 5466, 9542, 6956, 9980, 9868, 1194, 8239, 7591, 8296, 7511, 8556, 968, 9019, 7175, 4539, 5715, 9802, 4308, 7791, 3969, 4691, 7851, 5005, 8646, 3471, 1780, 9264, 7871, 2232, 8706, 1643, 6731, 7099, 4343, 6186, 4830, 6294, 1712, 8839, 1032, 7981, 7445, 9721, 1375, 5779, 3523, 1532, 2483, 7353, 8011, 8095, 7001, 9582, 3944, 8422, 9288, 3474, 7926, 8523, 3458, 1399, 1031, 9393, 2555, 3821, 211, 8539, 945, 1489, 104, 71, 2575, 1820, 8025, 2940, 8940, 411, 1166, 2292, 2743, 3030, 110, 8694, 4606, 7642, 5953, 1804, 377, 7586, 2100, 6183, 8540, 4819, 1111, 6867, 4298, 3583, 1532, 1263, 4908, 3188, 5951, 7922, 3972, 2506, 1660, 6198, 7340, 4272, 4429, 9852, 8250, 2135, 9354, 6002, 7300, 8110, 9035, 4922, 9932, 3853, 5762, 6951, 4690, 3344, 9429, 6634, 5416, 7128, 4005, 7182, 1659, 823, 3937, 5199, 6261, 1009, 8462, 7106, 2358, 3074, 6737, 7794, 6821, 5598, 162, 2683, 4664, 5312, 2903, 9012, 275, 6105, 721, 294, 6638, 9918, 2117, 5020, 897, 9615, 7435, 4081, 1587, 9052, 6106, 1222, 9442, 6904, 9288, 9388, 1901, 794, 4523, 2234, 3127, 3216, 3298, 5600, 2818, 48, 8553, 9439, 5282, 8589, 1548, 588, 3872, 9398, 2743, 398, 8254, 2379, 8086, 7657, 1157, 3761, 9079, 5525, 962, 5764, 7853, 7973, 9545, 7061, 1906, 7612, 1056, 6681, 6409, 809, 5289, 1419, 3421, 4161, 9171, 6377, 1749, 3346, 2535, 8219, 7993, 9654, 7337, 3675, 37, 4362, 3163, 4488, 9767, 4706, 2663, 6645, 8956, 7292, 5170, 4454, 7372, 5071, 1609}))
	//fmt.Println(minimumCardPickup([]int{1, 0, 5, 3}))
}
