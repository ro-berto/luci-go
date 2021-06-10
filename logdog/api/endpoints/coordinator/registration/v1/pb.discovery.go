// Code generated by cproto. DO NOT EDIT.

package logdog

import "go.chromium.org/luci/grpc/discovery"

import "google.golang.org/protobuf/types/descriptorpb"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"logdog.Registration",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 172, 121, 77, 143, 27, 71,
			122, 48, 201, 158, 25, 143, 202, 178, 37, 183, 228, 133, 204, 125,
			37, 63, 75, 127, 141, 36, 138, 51, 35, 189, 214, 198, 146, 199,
			8, 135, 108, 105, 218, 75, 145, 76, 55, 71, 178, 124, 177, 155,
			221, 69, 178, 188, 205, 170, 78, 85, 113, 70, 179, 73, 128, 0,
			251, 11, 178, 63, 96, 17, 228, 184, 200, 41, 64, 206, 73, 238,
			193, 34, 167, 28, 130, 252, 128, 252, 129, 4, 57, 6, 79, 85,
			55, 217, 28, 141, 118, 19, 36, 4, 52, 170, 207, 231, 187, 158,
			175, 38, 255, 244, 9, 185, 53, 21, 98, 154, 210, 221, 76, 10,
			45, 198, 139, 201, 110, 178, 144, 145, 102, 130, 183, 204, 138, 123,
			197, 238, 183, 138, 253, 198, 35, 178, 221, 205, 143, 184, 55, 200,
			91, 138, 198, 130, 39, 234, 70, 21, 170, 59, 78, 80, 76, 221,
			235, 100, 147, 71, 92, 168, 27, 53, 168, 238, 108, 6, 118, 114,
			248, 167, 228, 90, 44, 230, 173, 115, 32, 15, 223, 41, 0, 14,
			113, 101, 88, 253, 246, 206, 148, 233, 217, 98, 220, 138, 197, 124,
			119, 42, 210, 136, 79, 87, 244, 101, 250, 44, 163, 106, 73, 230,
			127, 86, 171, 127, 85, 115, 158, 14, 15, 127, 83, 187, 245, 212,
			194, 29, 230, 71, 91, 47, 104, 154, 254, 140, 139, 83, 62, 194,
			43, 95, 255, 203, 71, 100, 203, 221, 184, 85, 81, 85, 242, 143,
			151, 73, 245, 178, 235, 220, 170, 184, 247, 255, 225, 50, 152, 11,
			177, 72, 225, 112, 49, 153, 80, 169, 224, 30, 88, 80, 159, 41,
			72, 34, 29, 1, 227, 154, 202, 120, 22, 241, 41, 133, 137, 144,
			243, 72, 19, 232, 136, 236, 76, 178, 233, 76, 195, 253, 189, 189,
			63, 200, 47, 128, 207, 227, 22, 64, 59, 77, 193, 236, 41, 144,
			84, 81, 121, 66, 147, 22, 129, 153, 214, 153, 122, 180, 187, 155,
			208, 19, 154, 138, 140, 74, 85, 72, 2, 249, 204, 114, 34, 238,
			141, 45, 17, 187, 132, 64, 64, 19, 166, 180, 100, 227, 5, 178,
			10, 17, 79, 96, 161, 40, 48, 14, 74, 44, 100, 76, 205, 202,
			152, 241, 72, 158, 25, 186, 84, 19, 78, 153, 158, 129, 144, 230,
			127, 177, 208, 4, 230, 34, 97, 19, 22, 27, 97, 53, 33, 146,
			20, 50, 42, 231, 76, 107, 154, 64, 38, 197, 9, 75, 104, 2,
			122, 22, 105, 208, 51, 228, 46, 77, 197, 41, 227, 83, 64, 61,
			50, 188, 164, 240, 18, 129, 57, 213, 143, 8, 1, 252, 221, 57,
			71, 152, 2, 49, 41, 40, 138, 69, 66, 97, 190, 80, 26, 36,
			213, 17, 227, 6, 106, 52, 22, 39, 184, 149, 75, 140, 0, 23,
			154, 197, 180, 9, 122, 198, 20, 164, 76, 105, 132, 80, 198, 200,
			147, 115, 228, 36, 76, 197, 105, 196, 230, 84, 182, 222, 68, 4,
			227, 101, 89, 20, 68, 100, 82, 36, 139, 152, 174, 232, 32, 43,
			66, 254, 87, 116, 16, 200, 185, 75, 68, 188, 152, 83, 174, 163,
			66, 73, 187, 66, 130, 208, 51, 42, 97, 30, 105, 42, 89, 148,
			170, 149, 168, 141, 130, 244, 140, 18, 40, 83, 191, 100, 170, 79,
			153, 185, 137, 128, 121, 52, 167, 72, 80, 217, 182, 184, 88, 237,
			25, 185, 51, 173, 144, 35, 110, 65, 9, 169, 96, 30, 157, 193,
			152, 162, 165, 36, 160, 5, 80, 158, 8, 169, 40, 26, 69, 38,
			197, 92, 104, 10, 86, 38, 90, 65, 66, 37, 59, 161, 9, 76,
			164, 152, 19, 43, 5, 37, 38, 250, 20, 205, 36, 183, 32, 80,
			25, 141, 209, 130, 32, 147, 12, 13, 75, 162, 237, 112, 107, 69,
			74, 25, 218, 9, 140, 142, 252, 16, 194, 193, 147, 209, 139, 118,
			224, 129, 31, 194, 48, 24, 60, 247, 187, 94, 23, 14, 95, 194,
			232, 200, 131, 206, 96, 248, 50, 240, 159, 30, 141, 224, 104, 208,
			235, 122, 65, 8, 237, 126, 23, 58, 131, 254, 40, 240, 15, 143,
			71, 131, 32, 36, 208, 104, 135, 224, 135, 13, 179, 211, 238, 191,
			4, 239, 155, 97, 224, 133, 33, 12, 2, 240, 159, 13, 123, 190,
			215, 133, 23, 237, 32, 104, 247, 71, 190, 23, 54, 193, 239, 119,
			122, 199, 93, 191, 255, 180, 9, 135, 199, 35, 232, 15, 70, 4,
			122, 254, 51, 127, 228, 117, 97, 52, 104, 26, 180, 175, 223, 131,
			193, 19, 120, 230, 5, 157, 163, 118, 127, 212, 62, 244, 123, 254,
			232, 165, 65, 248, 196, 31, 245, 17, 217, 147, 65, 64, 160, 13,
			195, 118, 48, 242, 59, 199, 189, 118, 0, 195, 227, 96, 56, 8,
			61, 64, 206, 186, 126, 216, 233, 181, 253, 103, 94, 183, 5, 126,
			31, 250, 3, 240, 158, 123, 253, 17, 132, 71, 237, 94, 111, 157,
			81, 2, 131, 23, 125, 47, 64, 234, 203, 108, 194, 161, 7, 61,
			191, 125, 216, 243, 16, 149, 225, 179, 235, 7, 94, 103, 132, 12,
			173, 70, 29, 191, 235, 245, 71, 237, 94, 147, 64, 56, 244, 58,
			126, 187, 215, 4, 239, 27, 239, 217, 176, 215, 14, 94, 54, 115,
			160, 161, 247, 71, 199, 94, 127, 228, 183, 123, 208, 109, 63, 107,
			63, 245, 66, 216, 249, 125, 82, 25, 6, 131, 206, 113, 224, 61,
			67, 170, 7, 79, 32, 60, 62, 12, 71, 254, 232, 120, 228, 193,
			211, 193, 160, 107, 132, 29, 122, 193, 115, 191, 227, 133, 143, 161,
			55, 8, 141, 192, 142, 67, 175, 73, 160, 219, 30, 181, 13, 234,
			97, 48, 120, 226, 143, 194, 199, 56, 62, 60, 14, 125, 35, 56,
			191, 63, 242, 130, 224, 120, 56, 242, 7, 253, 219, 112, 52, 120,
			225, 61, 247, 2, 232, 180, 143, 67, 175, 107, 36, 60, 232, 35,
			183, 104, 43, 222, 32, 120, 137, 96, 81, 14, 70, 3, 77, 120,
			113, 228, 141, 142, 188, 0, 133, 106, 164, 213, 70, 49, 132, 163,
			192, 239, 140, 202, 199, 6, 1, 140, 6, 193, 136, 148, 248, 132,
			190, 247, 180, 231, 63, 245, 250, 29, 15, 183, 7, 8, 230, 133,
			31, 122, 183, 161, 29, 248, 33, 30, 240, 13, 98, 120, 209, 126,
			9, 131, 99, 195, 53, 42, 234, 56, 244, 136, 29, 151, 76, 183,
			105, 244, 9, 254, 19, 104, 119, 159, 251, 72, 121, 126, 122, 56,
			8, 67, 63, 55, 23, 35, 182, 206, 81, 46, 243, 22, 33, 219,
			164, 90, 115, 29, 168, 220, 192, 209, 182, 235, 52, 42, 143, 201,
			37, 82, 219, 254, 196, 14, 237, 226, 71, 149, 15, 205, 226, 135,
			118, 104, 23, 63, 174, 180, 205, 226, 219, 118, 104, 23, 63, 169,
			52, 205, 98, 213, 14, 237, 226, 167, 149, 150, 89, 204, 135, 118,
			241, 179, 74, 195, 44, 18, 59, 180, 139, 59, 149, 159, 152, 197,
			143, 237, 240, 215, 87, 73, 109, 163, 226, 110, 76, 42, 170, 90,
			255, 139, 171, 208, 134, 34, 220, 26, 239, 72, 21, 229, 90, 65,
			4, 138, 77, 57, 77, 154, 48, 97, 175, 104, 114, 47, 165, 124,
			170, 103, 160, 178, 136, 163, 151, 209, 108, 78, 87, 199, 105, 66,
			32, 194, 59, 177, 88, 112, 227, 51, 243, 184, 111, 28, 230, 68,
			70, 241, 42, 44, 20, 27, 26, 76, 14, 96, 166, 4, 195, 162,
			72, 173, 231, 3, 95, 3, 67, 239, 157, 208, 140, 242, 132, 90,
			128, 17, 63, 131, 56, 74, 41, 79, 34, 105, 160, 198, 130, 199,
			52, 211, 232, 166, 127, 78, 161, 145, 68, 103, 13, 130, 62, 173,
			49, 23, 92, 207, 26, 5, 24, 73, 211, 72, 91, 199, 55, 98,
			115, 170, 116, 52, 207, 172, 163, 206, 35, 92, 194, 48, 188, 82,
			30, 83, 24, 83, 125, 74, 41, 39, 160, 79, 203, 167, 79, 162,
			116, 65, 21, 2, 139, 86, 162, 66, 18, 152, 134, 56, 226, 232,
			91, 163, 4, 93, 185, 144, 160, 22, 99, 141, 236, 162, 68, 208,
			137, 66, 180, 2, 212, 130, 192, 100, 12, 8, 40, 203, 164, 120,
			197, 48, 28, 164, 103, 112, 247, 222, 254, 94, 115, 111, 111, 15,
			206, 104, 36, 21, 186, 207, 143, 192, 123, 21, 205, 179, 148, 42,
			66, 138, 33, 236, 63, 130, 142, 152, 103, 11, 77, 87, 100, 24,
			28, 107, 228, 154, 184, 151, 41, 186, 72, 132, 9, 190, 173, 60,
			72, 175, 248, 81, 58, 146, 26, 14, 160, 213, 106, 61, 62, 191,
			71, 121, 178, 182, 179, 68, 84, 228, 87, 197, 174, 221, 94, 38,
			135, 133, 90, 15, 16, 194, 114, 118, 207, 226, 42, 230, 143, 207,
			93, 50, 6, 144, 95, 177, 227, 226, 130, 153, 21, 72, 216, 4,
			118, 94, 67, 244, 37, 236, 193, 167, 159, 158, 135, 245, 21, 236,
			221, 134, 63, 177, 215, 46, 160, 238, 238, 1, 236, 63, 126, 109,
			55, 71, 125, 0, 251, 123, 197, 47, 63, 244, 103, 64, 83, 204,
			175, 46, 34, 224, 171, 11, 9, 248, 242, 119, 19, 112, 239, 119,
			16, 112, 247, 34, 2, 74, 234, 191, 191, 82, 255, 74, 95, 70,
			255, 171, 233, 221, 149, 194, 254, 231, 86, 240, 70, 93, 191, 217,
			70, 236, 86, 89, 229, 7, 235, 42, 135, 187, 175, 9, 225, 241,
			234, 82, 97, 0, 37, 165, 151, 47, 188, 102, 5, 171, 59, 235,
			114, 94, 179, 185, 178, 136, 87, 23, 46, 148, 238, 74, 189, 171,
			131, 95, 149, 15, 190, 1, 199, 221, 139, 113, 92, 104, 66, 37,
			13, 62, 120, 211, 3, 78, 34, 77, 209, 163, 182, 240, 79, 66,
			83, 83, 98, 192, 240, 76, 207, 108, 54, 133, 63, 141, 66, 127,
			253, 224, 78, 18, 157, 169, 131, 7, 77, 152, 51, 190, 208, 84,
			29, 236, 239, 221, 94, 127, 102, 112, 176, 196, 182, 115, 110, 171,
			245, 68, 138, 249, 104, 9, 74, 39, 183, 141, 239, 249, 58, 28,
			244, 225, 89, 148, 101, 140, 79, 9, 1, 159, 219, 21, 91, 233,
			52, 141, 215, 92, 210, 143, 213, 23, 122, 52, 202, 209, 204, 18,
			27, 6, 48, 143, 229, 83, 144, 81, 158, 186, 70, 232, 47, 9,
			136, 241, 15, 52, 214, 77, 56, 157, 81, 105, 19, 240, 252, 32,
			69, 161, 230, 217, 179, 90, 76, 38, 236, 21, 52, 84, 3, 118,
			24, 79, 76, 165, 194, 167, 69, 220, 184, 141, 142, 151, 32, 194,
			76, 210, 152, 34, 198, 241, 153, 77, 128, 23, 243, 49, 149, 165,
			16, 147, 215, 62, 171, 40, 163, 128, 190, 194, 128, 165, 12, 153,
			100, 25, 151, 162, 180, 184, 210, 130, 39, 66, 2, 181, 234, 106,
			194, 131, 101, 180, 50, 144, 246, 214, 96, 169, 153, 88, 164, 9,
			140, 41, 89, 242, 206, 214, 4, 133, 162, 104, 60, 80, 13, 228,
			151, 161, 246, 215, 130, 226, 126, 9, 88, 14, 139, 96, 36, 89,
			145, 120, 17, 180, 86, 97, 93, 251, 8, 23, 225, 156, 131, 74,
			96, 206, 98, 185, 14, 247, 191, 11, 118, 95, 53, 90, 196, 252,
			156, 141, 74, 213, 117, 38, 219, 87, 201, 191, 86, 201, 198, 70,
			165, 86, 113, 157, 31, 106, 215, 235, 191, 173, 66, 104, 178, 130,
			37, 82, 76, 5, 80, 105, 165, 180, 160, 5, 207, 176, 210, 26,
			83, 107, 219, 247, 30, 236, 127, 222, 252, 252, 167, 15, 49, 192,
			225, 63, 130, 161, 248, 238, 185, 69, 96, 60, 78, 23, 138, 157,
			208, 22, 244, 133, 166, 143, 16, 170, 162, 48, 22, 11, 195, 154,
			196, 106, 209, 188, 28, 91, 155, 60, 34, 240, 112, 15, 137, 216,
			157, 51, 14, 119, 112, 50, 103, 124, 119, 38, 225, 14, 220, 255,
			255, 48, 147, 187, 73, 116, 6, 119, 224, 193, 195, 207, 91, 247,
			63, 7, 124, 35, 187, 24, 92, 225, 142, 125, 161, 54, 210, 18,
			114, 153, 108, 34, 119, 155, 200, 222, 91, 197, 172, 234, 58, 63,
			108, 95, 41, 102, 142, 235, 252, 224, 94, 35, 191, 116, 140, 32,
			170, 174, 35, 107, 110, 253, 223, 107, 133, 32, 214, 146, 155, 40,
			151, 203, 122, 118, 83, 74, 110, 202, 242, 34, 43, 129, 21, 175,
			73, 65, 74, 149, 178, 15, 70, 112, 186, 132, 38, 215, 114, 45,
			107, 141, 17, 236, 17, 248, 62, 215, 195, 247, 48, 97, 52, 77,
			140, 13, 68, 144, 9, 197, 52, 59, 49, 37, 30, 167, 211, 200,
			140, 191, 55, 4, 229, 7, 173, 161, 23, 110, 64, 25, 82, 74,
			8, 133, 132, 185, 144, 180, 9, 17, 112, 193, 239, 253, 130, 74,
			97, 179, 32, 180, 26, 195, 192, 26, 52, 91, 90, 227, 75, 40,
			216, 195, 66, 21, 243, 71, 52, 47, 115, 124, 157, 206, 243, 38,
			242, 197, 23, 95, 52, 243, 127, 214, 60, 74, 11, 37, 211, 40,
			244, 85, 221, 68, 45, 20, 250, 170, 162, 78, 182, 223, 41, 102,
			142, 235, 200, 171, 239, 141, 183, 76, 255, 228, 1, 249, 243, 15,
			73, 56, 21, 173, 120, 38, 197, 156, 45, 230, 45, 33, 167, 187,
			233, 34, 102, 187, 169, 152, 38, 98, 186, 27, 101, 108, 151, 242,
			36, 19, 140, 107, 181, 27, 11, 33, 19, 198, 35, 45, 228, 174,
			164, 83, 44, 200, 141, 132, 118, 79, 246, 119, 21, 149, 39, 44,
			206, 219, 84, 238, 150, 189, 95, 255, 61, 253, 178, 198, 63, 87,
			201, 251, 129, 129, 68, 229, 80, 210, 9, 123, 21, 208, 63, 94,
			80, 165, 221, 27, 228, 173, 76, 10, 244, 136, 166, 89, 118, 41,
			40, 166, 238, 143, 200, 86, 102, 142, 154, 110, 217, 165, 32, 159,
			185, 215, 201, 166, 164, 81, 58, 191, 177, 105, 150, 237, 196, 253,
			144, 188, 109, 123, 45, 223, 49, 62, 17, 55, 28, 112, 118, 46,
			5, 196, 46, 249, 124, 34, 220, 15, 200, 182, 200, 190, 227, 152,
			47, 223, 216, 128, 234, 206, 229, 224, 45, 145, 245, 113, 234, 126,
			65, 8, 125, 149, 49, 75, 241, 13, 2, 213, 157, 183, 239, 127,
			112, 190, 31, 215, 42, 204, 51, 40, 29, 110, 124, 75, 126, 116,
			158, 47, 149, 9, 174, 40, 146, 175, 104, 44, 169, 229, 235, 114,
			144, 207, 220, 29, 114, 53, 21, 211, 239, 198, 11, 158, 164, 244,
			59, 45, 50, 22, 231, 12, 190, 155, 138, 233, 161, 89, 30, 225,
			234, 253, 239, 200, 229, 160, 36, 125, 119, 64, 222, 93, 199, 229,
			222, 108, 89, 249, 183, 46, 148, 109, 253, 214, 155, 182, 45, 137,
			135, 253, 111, 123, 255, 39, 70, 241, 216, 94, 248, 250, 87, 117,
			178, 229, 110, 108, 84, 70, 85, 242, 183, 85, 211, 74, 220, 168,
			184, 247, 127, 83, 93, 235, 10, 238, 63, 132, 209, 140, 66, 239,
			184, 227, 67, 123, 161, 103, 66, 170, 214, 27, 90, 131, 199, 138,
			218, 183, 100, 26, 48, 171, 70, 26, 83, 48, 21, 39, 84, 162,
			231, 89, 240, 36, 239, 11, 181, 179, 40, 70, 192, 44, 166, 92,
			209, 38, 60, 167, 82, 161, 179, 185, 223, 66, 111, 139, 133, 78,
			94, 166, 76, 208, 157, 22, 129, 182, 231, 119, 188, 126, 232, 193,
			132, 165, 116, 89, 179, 110, 85, 174, 228, 85, 227, 118, 37, 46,
			42, 209, 124, 232, 84, 92, 135, 84, 118, 200, 208, 214, 143, 239,
			84, 126, 90, 173, 119, 225, 66, 5, 128, 204, 87, 49, 27, 224,
			244, 20, 236, 238, 178, 217, 5, 157, 149, 76, 91, 171, 136, 243,
			206, 246, 77, 178, 83, 4, 156, 43, 181, 247, 235, 63, 54, 18,
			75, 197, 20, 51, 5, 26, 205, 63, 51, 141, 51, 124, 36, 173,
			178, 243, 190, 82, 219, 46, 57, 239, 43, 151, 174, 150, 156, 247,
			149, 107, 215, 201, 126, 225, 187, 221, 218, 245, 250, 199, 231, 96,
			130, 125, 93, 232, 119, 10, 170, 215, 60, 141, 187, 4, 142, 158,
			198, 189, 116, 165, 228, 105, 92, 247, 26, 249, 15, 27, 34, 107,
			174, 83, 175, 93, 171, 255, 91, 213, 128, 55, 79, 211, 54, 236,
			118, 144, 233, 92, 230, 57, 241, 183, 17, 89, 164, 148, 136, 89,
			164, 151, 153, 80, 137, 22, 188, 98, 218, 104, 51, 83, 194, 34,
			172, 88, 112, 29, 49, 174, 160, 221, 233, 41, 72, 232, 132, 113,
			204, 139, 78, 103, 2, 78, 89, 154, 154, 42, 116, 156, 82, 203,
			70, 148, 32, 131, 106, 105, 35, 76, 145, 28, 180, 133, 74, 177,
			150, 78, 169, 44, 50, 132, 89, 116, 66, 161, 145, 191, 25, 188,
			217, 138, 37, 141, 52, 109, 148, 218, 122, 214, 110, 10, 122, 150,
			34, 170, 109, 34, 227, 133, 136, 80, 200, 245, 75, 239, 22, 51,
			199, 117, 234, 239, 185, 228, 15, 141, 132, 28, 215, 185, 89, 107,
			212, 31, 192, 32, 203, 211, 46, 116, 88, 152, 133, 152, 130, 122,
			44, 22, 182, 34, 47, 180, 128, 236, 69, 83, 202, 87, 186, 118,
			54, 16, 196, 114, 182, 233, 58, 55, 223, 126, 175, 152, 85, 93,
			231, 166, 123, 179, 152, 33, 50, 248, 9, 249, 235, 77, 131, 122,
			195, 117, 154, 181, 247, 235, 191, 222, 92, 225, 54, 254, 208, 232,
			33, 77, 197, 41, 72, 170, 37, 163, 106, 249, 234, 130, 97, 167,
			5, 237, 94, 15, 58, 61, 223, 235, 143, 66, 8, 143, 6, 199,
			189, 110, 209, 199, 180, 29, 206, 150, 21, 101, 202, 40, 215, 133,
			40, 167, 148, 83, 89, 104, 213, 34, 49, 127, 108, 18, 152, 73,
			154, 69, 134, 51, 203, 169, 121, 45, 4, 230, 84, 169, 104, 74,
			155, 69, 131, 153, 131, 164, 247, 22, 138, 174, 162, 169, 133, 132,
			225, 119, 141, 210, 37, 140, 66, 175, 246, 220, 42, 241, 123, 112,
			31, 198, 103, 218, 30, 151, 17, 79, 108, 181, 17, 181, 202, 167,
			77, 248, 230, 194, 4, 100, 73, 77, 159, 56, 111, 133, 44, 187,
			35, 186, 192, 163, 96, 71, 240, 244, 172, 76, 9, 214, 20, 229,
			192, 159, 159, 188, 141, 36, 245, 7, 35, 239, 145, 53, 99, 166,
			32, 94, 72, 132, 149, 158, 129, 200, 213, 208, 132, 241, 66, 91,
			1, 69, 105, 154, 139, 82, 89, 123, 92, 100, 83, 25, 153, 175,
			17, 162, 232, 69, 199, 51, 58, 167, 77, 96, 186, 176, 249, 88,
			204, 41, 204, 35, 158, 160, 39, 57, 51, 153, 84, 33, 92, 45,
			35, 174, 152, 173, 133, 115, 165, 226, 51, 16, 246, 91, 71, 211,
			138, 185, 63, 128, 192, 27, 5, 190, 23, 174, 94, 17, 218, 3,
			102, 118, 66, 22, 246, 110, 248, 105, 90, 48, 152, 36, 201, 19,
			42, 241, 65, 137, 24, 115, 106, 101, 26, 65, 66, 74, 26, 35,
			107, 227, 220, 140, 115, 179, 152, 68, 44, 85, 104, 103, 83, 90,
			152, 183, 13, 66, 121, 239, 102, 70, 73, 14, 113, 105, 230, 27,
			155, 104, 174, 69, 126, 179, 81, 117, 157, 230, 118, 225, 210, 54,
			28, 215, 105, 94, 187, 78, 254, 210, 230, 163, 155, 174, 243, 176,
			118, 183, 254, 43, 199, 168, 51, 247, 30, 171, 40, 157, 231, 151,
			254, 4, 190, 60, 128, 189, 102, 217, 7, 125, 102, 156, 72, 180,
			72, 117, 126, 141, 148, 239, 101, 84, 50, 145, 172, 100, 146, 101,
			41, 195, 200, 68, 202, 120, 138, 221, 56, 21, 106, 85, 136, 149,
			92, 59, 68, 19, 77, 37, 48, 173, 94, 135, 221, 34, 48, 64,
			181, 219, 203, 77, 19, 37, 114, 39, 88, 14, 178, 43, 187, 51,
			216, 184, 128, 84, 240, 41, 149, 134, 170, 56, 166, 153, 182, 100,
			249, 19, 80, 11, 75, 101, 254, 249, 198, 38, 172, 249, 75, 160,
			175, 98, 74, 237, 183, 27, 20, 9, 186, 154, 220, 100, 83, 17,
			71, 41, 232, 72, 253, 188, 105, 218, 1, 4, 148, 152, 211, 98,
			23, 215, 63, 83, 57, 97, 203, 79, 41, 137, 20, 89, 70, 19,
			72, 22, 198, 131, 160, 138, 77, 189, 91, 162, 123, 169, 204, 205,
			45, 84, 81, 189, 152, 85, 93, 231, 225, 143, 63, 41, 102, 142,
			235, 60, 220, 185, 67, 190, 34, 168, 229, 141, 71, 149, 195, 106,
			253, 126, 30, 63, 114, 35, 201, 61, 195, 50, 241, 94, 15, 183,
			198, 77, 217, 248, 137, 193, 233, 209, 246, 45, 242, 75, 12, 71,
			85, 12, 160, 7, 181, 107, 245, 19, 8, 77, 242, 133, 175, 79,
			47, 53, 135, 28, 153, 229, 150, 125, 154, 121, 254, 110, 83, 109,
			124, 113, 39, 84, 142, 35, 205, 230, 232, 243, 15, 23, 58, 69,
			115, 183, 121, 155, 202, 67, 23, 149, 26, 196, 41, 167, 82, 205,
			88, 182, 244, 154, 203, 24, 131, 252, 85, 77, 108, 62, 200, 13,
			185, 106, 98, 243, 193, 246, 187, 197, 204, 113, 157, 131, 247, 92,
			66, 13, 185, 85, 215, 105, 215, 110, 213, 191, 177, 94, 41, 90,
			41, 96, 184, 24, 239, 134, 139, 49, 152, 116, 17, 113, 103, 139,
			113, 202, 212, 12, 159, 89, 74, 165, 201, 81, 239, 217, 40, 130,
			165, 145, 37, 22, 131, 95, 153, 224, 37, 65, 24, 207, 219, 121,
			176, 170, 154, 120, 222, 190, 244, 65, 49, 115, 92, 167, 253, 255,
			110, 146, 191, 171, 146, 218, 86, 197, 221, 120, 90, 25, 85, 235,
			127, 83, 133, 114, 38, 10, 121, 17, 96, 155, 193, 61, 49, 237,
			138, 233, 154, 193, 23, 89, 163, 77, 187, 204, 119, 226, 73, 20,
			83, 219, 65, 32, 197, 141, 156, 74, 198, 149, 142, 56, 238, 162,
			219, 167, 28, 227, 183, 58, 151, 9, 172, 189, 5, 60, 86, 168,
			195, 30, 67, 179, 99, 156, 105, 22, 165, 236, 23, 133, 225, 17,
			226, 108, 161, 176, 159, 110, 95, 39, 191, 221, 32, 27, 91, 38,
			159, 10, 107, 207, 234, 127, 191, 177, 52, 160, 208, 162, 49, 222,
			14, 153, 57, 71, 83, 57, 29, 130, 168, 156, 46, 173, 62, 88,
			150, 147, 56, 56, 206, 80, 60, 139, 24, 93, 98, 243, 53, 71,
			96, 30, 175, 164, 122, 33, 249, 26, 75, 100, 61, 7, 224, 73,
			137, 175, 44, 146, 209, 156, 154, 36, 82, 11, 3, 209, 210, 184,
			76, 140, 86, 81, 206, 100, 51, 198, 3, 93, 200, 200, 116, 245,
			225, 153, 165, 76, 159, 149, 236, 200, 212, 155, 81, 33, 109, 149,
			69, 49, 109, 149, 19, 36, 243, 52, 76, 56, 122, 17, 248, 35,
			207, 248, 28, 101, 8, 66, 175, 182, 244, 165, 133, 62, 10, 8,
			254, 100, 253, 60, 83, 38, 192, 230, 53, 124, 238, 161, 140, 80,
			208, 121, 148, 100, 218, 24, 46, 19, 174, 46, 229, 140, 38, 13,
			152, 6, 195, 206, 178, 137, 219, 62, 103, 31, 185, 71, 42, 84,
			69, 77, 239, 97, 46, 148, 13, 170, 45, 104, 39, 246, 219, 117,
			148, 158, 147, 251, 186, 95, 93, 163, 2, 26, 237, 20, 115, 200,
			51, 239, 21, 83, 90, 173, 83, 112, 153, 108, 110, 217, 60, 59,
			220, 186, 94, 204, 106, 174, 19, 190, 127, 167, 152, 57, 174, 19,
			126, 254, 179, 162, 4, 255, 175, 0, 0, 0, 255, 255, 215, 31,
			201, 209, 86, 34, 0, 0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptorpb.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("logdog.Registration")
	if err != nil {
		panic(err)
	}
	return ret
}
