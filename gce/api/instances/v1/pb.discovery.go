// Code generated by cproto. DO NOT EDIT.

package instances

import "go.chromium.org/luci/grpc/discovery"

import "google.golang.org/protobuf/types/descriptorpb"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"instances.Instances",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 236, 123, 77, 108, 27, 199,
			150, 46, 187, 155, 162, 165, 146, 127, 219, 178, 163, 48, 137, 115,
			34, 71, 49, 153, 75, 145, 250, 177, 227, 216, 78, 94, 64, 73,
			148, 212, 142, 76, 42, 36, 101, 199, 14, 242, 236, 86, 119, 145,
			236, 184, 217, 197, 219, 213, 148, 172, 56, 6, 238, 187, 120, 239,
			102, 241, 128, 139, 247, 102, 63, 192, 96, 150, 119, 49, 155, 139,
			89, 15, 102, 49, 152, 205, 0, 179, 187, 203, 217, 205, 122, 54,
			179, 24, 96, 128, 193, 224, 156, 234, 110, 82, 182, 252, 147, 12,
			48, 171, 8, 22, 212, 167, 186, 234, 252, 213, 169, 115, 190, 170,
			106, 179, 191, 49, 217, 59, 93, 33, 186, 62, 175, 12, 66, 17,
			137, 189, 97, 167, 194, 251, 131, 232, 176, 76, 164, 121, 70, 189,
			44, 39, 47, 231, 78, 176, 137, 26, 190, 95, 125, 198, 206, 59,
			162, 95, 126, 238, 253, 42, 163, 183, 59, 72, 238, 104, 15, 146,
			215, 93, 225, 219, 65, 183, 44, 194, 238, 72, 76, 116, 56, 224,
			178, 242, 56, 16, 7, 129, 18, 57, 216, 251, 55, 77, 251, 75,
			221, 216, 220, 89, 253, 131, 126, 105, 83, 141, 220, 137, 187, 151,
			239, 113, 223, 255, 18, 59, 183, 113, 220, 237, 255, 56, 203, 114,
			102, 246, 82, 102, 229, 44, 251, 135, 147, 76, 59, 105, 26, 151,
			50, 230, 242, 223, 158, 4, 26, 224, 8, 31, 86, 135, 157, 14,
			15, 37, 44, 128, 98, 117, 69, 130, 107, 71, 54, 120, 65, 196,
			67, 167, 103, 7, 93, 14, 29, 17, 246, 237, 136, 193, 154, 24,
			28, 134, 94, 183, 23, 193, 242, 226, 226, 167, 241, 0, 176, 2,
			167, 12, 80, 245, 125, 160, 119, 18, 66, 46, 121, 184, 207, 221,
			50, 131, 94, 20, 13, 228, 205, 74, 197, 229, 251, 220, 23, 3,
			30, 202, 196, 86, 71, 244, 149, 145, 142, 240, 23, 246, 148, 18,
			21, 198, 160, 201, 93, 79, 70, 161, 183, 55, 140, 60, 17, 128,
			29, 184, 48, 148, 28, 188, 0, 164, 24, 134, 14, 167, 150, 61,
			47, 176, 195, 67, 210, 75, 150, 224, 192, 139, 122, 32, 66, 250,
			43, 134, 17, 131, 190, 112, 189, 142, 231, 216, 200, 161, 4, 118,
			200, 97, 192, 195, 190, 23, 69, 220, 133, 65, 40, 246, 61, 151,
			187, 16, 245, 236, 8, 162, 30, 90, 231, 251, 226, 192, 11, 186,
			224, 136, 192, 245, 112, 144, 196, 65, 12, 250, 60, 186, 201, 24,
			224, 207, 199, 207, 41, 38, 65, 116, 18, 141, 28, 225, 114, 232,
			15, 101, 4, 33, 143, 108, 47, 32, 174, 246, 158, 216, 199, 87,
			177, 199, 24, 4, 34, 242, 28, 94, 130, 168, 231, 73, 240, 61,
			25, 33, 135, 113, 137, 129, 251, 156, 58, 174, 39, 29, 223, 246,
			250, 60, 44, 191, 76, 9, 47, 24, 247, 69, 162, 196, 32, 20,
			238, 208, 225, 35, 61, 216, 72, 145, 255, 146, 30, 12, 98, 235,
			92, 225, 12, 251, 60, 136, 236, 100, 146, 42, 34, 4, 17, 245,
			120, 8, 125, 59, 226, 161, 103, 251, 114, 228, 106, 154, 160, 168,
			199, 25, 140, 107, 159, 26, 85, 231, 30, 141, 68, 198, 129, 221,
			231, 168, 208, 120, 108, 5, 98, 244, 142, 252, 238, 69, 18, 45,
			10, 20, 43, 17, 74, 232, 219, 135, 176, 199, 49, 82, 92, 136,
			4, 240, 192, 21, 161, 228, 24, 20, 131, 80, 244, 69, 196, 65,
			249, 36, 146, 224, 242, 208, 219, 231, 46, 116, 66, 209, 103, 202,
			11, 82, 116, 162, 3, 12, 147, 56, 130, 64, 14, 184, 131, 17,
			4, 131, 208, 195, 192, 10, 49, 118, 2, 21, 69, 82, 146, 238,
			12, 218, 91, 86, 11, 90, 141, 141, 246, 189, 106, 179, 6, 86,
			11, 118, 154, 141, 187, 214, 122, 109, 29, 86, 239, 67, 123, 171,
			6, 107, 141, 157, 251, 77, 107, 115, 171, 13, 91, 141, 237, 245,
			90, 179, 5, 213, 250, 58, 172, 53, 234, 237, 166, 181, 186, 219,
			110, 52, 91, 12, 230, 170, 45, 176, 90, 115, 244, 166, 90, 191,
			15, 181, 175, 119, 154, 181, 86, 11, 26, 77, 176, 238, 236, 108,
			91, 181, 117, 184, 87, 109, 54, 171, 245, 182, 85, 107, 149, 192,
			170, 175, 109, 239, 174, 91, 245, 205, 18, 172, 238, 182, 161, 222,
			104, 51, 216, 182, 238, 88, 237, 218, 58, 180, 27, 37, 18, 251,
			226, 56, 104, 108, 192, 157, 90, 115, 109, 171, 90, 111, 87, 87,
			173, 109, 171, 125, 159, 4, 110, 88, 237, 58, 10, 219, 104, 52,
			25, 84, 97, 167, 218, 108, 91, 107, 187, 219, 213, 38, 236, 236,
			54, 119, 26, 173, 26, 160, 101, 235, 86, 107, 109, 187, 106, 221,
			169, 173, 151, 193, 170, 67, 189, 1, 181, 187, 181, 122, 27, 90,
			91, 213, 237, 237, 163, 134, 50, 104, 220, 171, 215, 154, 168, 253,
			184, 153, 176, 90, 131, 109, 171, 186, 186, 93, 67, 81, 100, 231,
			186, 213, 172, 173, 181, 209, 160, 209, 211, 154, 181, 94, 171, 183,
			171, 219, 37, 6, 173, 157, 218, 154, 85, 221, 46, 65, 237, 235,
			218, 157, 157, 237, 106, 243, 126, 41, 102, 218, 170, 125, 181, 91,
			171, 183, 173, 234, 54, 172, 87, 239, 84, 55, 107, 45, 40, 188,
			206, 43, 59, 205, 198, 218, 110, 179, 118, 7, 181, 110, 108, 64,
			107, 119, 181, 213, 182, 218, 187, 237, 26, 108, 54, 26, 235, 228,
			236, 86, 173, 121, 215, 90, 171, 181, 110, 193, 118, 163, 69, 14,
			219, 109, 213, 74, 12, 214, 171, 237, 42, 137, 222, 105, 54, 54,
			172, 118, 235, 22, 62, 175, 238, 182, 44, 114, 156, 85, 111, 215,
			154, 205, 221, 157, 182, 213, 168, 23, 97, 171, 113, 175, 118, 183,
			214, 132, 181, 234, 110, 171, 182, 78, 30, 110, 212, 209, 90, 140,
			149, 90, 163, 121, 31, 217, 162, 31, 104, 6, 74, 112, 111, 171,
			214, 222, 170, 53, 209, 169, 228, 173, 42, 186, 161, 213, 110, 90,
			107, 237, 241, 110, 141, 38, 180, 27, 205, 54, 27, 179, 19, 234,
			181, 205, 109, 107, 179, 86, 95, 171, 225, 235, 6, 178, 185, 103,
			181, 106, 69, 168, 54, 173, 22, 118, 176, 72, 48, 220, 171, 222,
			135, 198, 46, 89, 141, 19, 181, 219, 170, 49, 245, 60, 22, 186,
			37, 154, 79, 176, 54, 160, 186, 126, 215, 66, 205, 227, 222, 59,
			141, 86, 203, 138, 195, 133, 220, 182, 182, 21, 251, 188, 204, 216,
			36, 211, 116, 211, 128, 204, 44, 62, 77, 154, 198, 92, 230, 22,
			155, 98, 250, 228, 188, 122, 84, 141, 151, 51, 53, 106, 156, 86,
			143, 170, 241, 195, 76, 137, 26, 53, 245, 168, 26, 231, 51, 191,
			162, 198, 248, 81, 53, 126, 148, 153, 163, 70, 166, 30, 85, 227,
			149, 204, 7, 212, 248, 161, 122, 84, 141, 133, 204, 251, 212, 248,
			190, 122, 252, 119, 157, 233, 217, 140, 105, 172, 100, 206, 230, 255,
			69, 135, 42, 116, 121, 192, 67, 207, 1, 170, 160, 208, 231, 82,
			218, 93, 174, 74, 192, 161, 24, 130, 99, 7, 16, 242, 5, 44,
			52, 145, 0, 123, 95, 120, 46, 184, 188, 227, 5, 148, 254, 134,
			3, 31, 139, 9, 119, 217, 209, 241, 148, 126, 15, 197, 48, 132,
			234, 142, 37, 203, 80, 133, 232, 112, 224, 57, 182, 15, 252, 137,
			221, 31, 248, 28, 60, 137, 252, 168, 126, 69, 96, 75, 202, 98,
			33, 255, 245, 144, 203, 136, 65, 156, 213, 66, 46, 7, 34, 64,
			201, 135, 3, 74, 125, 118, 128, 252, 176, 248, 244, 132, 91, 134,
			13, 17, 130, 23, 200, 200, 14, 28, 158, 84, 35, 172, 175, 158,
			195, 97, 67, 8, 120, 170, 154, 0, 194, 129, 3, 171, 118, 88,
			120, 14, 107, 148, 9, 106, 20, 177, 54, 13, 195, 64, 194, 75,
			222, 223, 82, 108, 158, 97, 98, 235, 113, 184, 221, 106, 212, 169,
			146, 112, 153, 166, 249, 142, 8, 225, 17, 245, 126, 132, 150, 41,
			95, 80, 71, 177, 247, 29, 119, 34, 120, 244, 244, 217, 163, 50,
			99, 140, 25, 217, 140, 102, 26, 43, 147, 167, 246, 114, 36, 102,
			133, 253, 253, 18, 123, 255, 121, 4, 21, 121, 125, 46, 35, 187,
			63, 120, 25, 138, 186, 197, 166, 218, 73, 31, 115, 150, 157, 144,
			28, 235, 148, 156, 213, 64, 43, 24, 205, 132, 52, 103, 216, 68,
			96, 7, 66, 206, 234, 160, 21, 38, 154, 138, 88, 253, 63, 218,
			241, 208, 235, 116, 202, 50, 129, 95, 203, 111, 8, 191, 82, 125,
			127, 18, 4, 251, 187, 10, 59, 97, 78, 92, 202, 252, 153, 166,
			253, 130, 193, 126, 193, 96, 191, 96, 176, 95, 48, 216, 47, 24,
			236, 23, 12, 246, 223, 136, 193, 82, 104, 132, 143, 9, 6, 179,
			18, 96, 134, 143, 9, 6, 75, 129, 217, 124, 10, 204, 62, 202,
			84, 18, 96, 134, 143, 9, 6, 75, 129, 217, 149, 20, 152, 21,
			70, 192, 12, 31, 255, 233, 61, 194, 96, 185, 31, 53, 44, 125,
			249, 127, 124, 15, 170, 144, 150, 222, 17, 180, 144, 96, 195, 64,
			120, 65, 68, 105, 205, 235, 99, 153, 113, 249, 128, 7, 46, 15,
			34, 5, 135, 14, 85, 251, 247, 34, 160, 108, 226, 11, 199, 246,
			25, 56, 182, 207, 3, 215, 14, 75, 192, 3, 204, 254, 46, 2,
			44, 27, 28, 49, 84, 227, 98, 116, 64, 185, 180, 19, 218, 206,
			168, 98, 36, 47, 176, 32, 32, 84, 32, 26, 43, 166, 240, 85,
			82, 36, 4, 164, 24, 121, 88, 74, 125, 59, 242, 246, 21, 52,
			12, 128, 15, 132, 211, 3, 59, 130, 221, 246, 26, 244, 61, 55,
			160, 140, 46, 2, 6, 183, 237, 96, 136, 101, 96, 169, 4, 75,
			55, 174, 47, 150, 146, 68, 61, 8, 133, 207, 7, 145, 231, 192,
			102, 200, 187, 34, 244, 236, 32, 213, 30, 14, 122, 158, 211, 3,
			254, 36, 226, 168, 19, 37, 232, 99, 122, 237, 217, 206, 227, 3,
			59, 116, 9, 79, 30, 114, 59, 4, 17, 112, 76, 128, 88, 242,
			251, 94, 48, 140, 56, 213, 75, 248, 100, 49, 181, 207, 23, 65,
			183, 12, 219, 220, 30, 140, 76, 14, 57, 204, 201, 62, 183, 67,
			238, 206, 129, 20, 170, 0, 7, 2, 124, 110, 15, 88, 220, 13,
			34, 123, 79, 97, 215, 128, 115, 244, 107, 135, 16, 104, 196, 195,
			1, 214, 86, 85, 208, 135, 18, 171, 146, 13, 223, 44, 95, 93,
			232, 33, 4, 246, 189, 128, 219, 33, 3, 226, 254, 109, 225, 213,
			160, 3, 231, 179, 66, 61, 139, 229, 24, 112, 134, 132, 114, 60,
			73, 53, 1, 22, 23, 23, 151, 22, 232, 95, 123, 113, 241, 38,
			253, 123, 128, 166, 223, 184, 113, 227, 198, 194, 210, 242, 194, 202,
			82, 123, 121, 229, 230, 181, 27, 55, 175, 221, 40, 223, 72, 126,
			30, 148, 97, 245, 144, 225, 68, 70, 161, 231, 68, 168, 96, 20,
			155, 72, 220, 75, 112, 192, 129, 7, 114, 24, 198, 208, 255, 128,
			19, 242, 119, 68, 176, 207, 195, 72, 205, 175, 42, 74, 240, 77,
			115, 99, 141, 193, 202, 202, 202, 141, 145, 45, 7, 7, 7, 101,
			143, 71, 29, 66, 136, 97, 199, 193, 95, 236, 81, 142, 158, 68,
			69, 68, 108, 28, 80, 114, 208, 149, 104, 212, 101, 168, 169, 93,
			128, 100, 44, 121, 132, 165, 155, 176, 38, 250, 131, 97, 196, 199,
			214, 2, 9, 220, 105, 180, 172, 175, 225, 17, 122, 166, 80, 68,
			20, 77, 133, 121, 212, 41, 5, 159, 49, 80, 31, 129, 103, 201,
			163, 135, 241, 4, 23, 104, 120, 125, 119, 123, 187, 88, 60, 182,
			31, 197, 123, 97, 177, 120, 107, 76, 167, 229, 215, 233, 212, 229,
			17, 114, 17, 29, 215, 62, 28, 211, 77, 70, 225, 208, 137, 72,
			192, 190, 237, 67, 180, 31, 75, 60, 210, 253, 163, 104, 191, 4,
			164, 208, 173, 159, 107, 210, 126, 57, 218, 71, 234, 85, 22, 169,
			78, 67, 201, 29, 248, 24, 150, 22, 23, 143, 90, 184, 242, 82,
			11, 239, 121, 193, 202, 50, 60, 218, 228, 81, 235, 80, 70, 188,
			143, 175, 171, 114, 195, 243, 121, 251, 232, 68, 108, 88, 219, 181,
			182, 117, 167, 6, 157, 40, 86, 227, 101, 99, 62, 234, 68, 137,
			166, 187, 86, 189, 253, 201, 85, 136, 60, 231, 177, 132, 207, 161,
			80, 40, 168, 150, 98, 39, 42, 187, 7, 91, 94, 183, 183, 110,
			71, 52, 170, 8, 159, 125, 6, 43, 203, 69, 248, 1, 232, 221,
			182, 56, 72, 94, 37, 126, 171, 84, 160, 138, 250, 186, 226, 64,
			18, 75, 92, 44, 75, 139, 139, 99, 57, 76, 150, 211, 14, 42,
			75, 45, 125, 242, 226, 50, 74, 185, 225, 240, 165, 79, 174, 94,
			189, 122, 125, 229, 147, 197, 81, 218, 216, 227, 29, 17, 114, 216,
			13, 188, 39, 9, 151, 27, 215, 23, 159, 231, 82, 254, 121, 147,
			89, 80, 246, 67, 161, 160, 156, 82, 161, 201, 194, 159, 34, 44,
			140, 171, 243, 154, 8, 70, 62, 232, 174, 132, 207, 252, 24, 31,
			10, 128, 226, 145, 0, 184, 250, 210, 0, 184, 109, 239, 219, 240,
			72, 77, 100, 217, 25, 134, 33, 15, 34, 236, 114, 199, 243, 125,
			79, 142, 5, 0, 102, 83, 232, 83, 43, 124, 14, 47, 31, 240,
			138, 48, 135, 207, 71, 173, 229, 128, 31, 172, 14, 61, 223, 229,
			97, 161, 136, 134, 181, 98, 15, 197, 34, 148, 99, 138, 201, 222,
			30, 0, 251, 212, 149, 237, 94, 16, 161, 229, 113, 79, 101, 122,
			108, 54, 121, 160, 88, 222, 67, 206, 164, 203, 200, 7, 215, 94,
			227, 3, 139, 206, 24, 162, 114, 32, 14, 198, 204, 142, 91, 33,
			16, 7, 240, 57, 28, 233, 243, 74, 75, 71, 138, 191, 222, 228,
			64, 28, 148, 187, 60, 170, 97, 176, 169, 182, 66, 113, 204, 242,
			163, 214, 199, 157, 145, 40, 188, 196, 210, 79, 94, 106, 105, 60,
			95, 9, 206, 128, 157, 195, 168, 167, 54, 18, 71, 2, 109, 124,
			162, 10, 197, 231, 163, 112, 147, 71, 107, 163, 121, 47, 20, 41,
			215, 211, 49, 200, 29, 123, 48, 240, 130, 46, 99, 96, 5, 170,
			69, 237, 218, 75, 4, 3, 198, 252, 116, 56, 160, 82, 119, 4,
			184, 168, 210, 17, 99, 6, 70, 5, 232, 39, 213, 31, 37, 10,
			177, 139, 141, 176, 165, 164, 216, 168, 86, 20, 54, 247, 20, 113,
			195, 179, 133, 167, 125, 17, 68, 189, 103, 11, 79, 93, 251, 240,
			89, 251, 41, 22, 239, 103, 55, 159, 246, 189, 224, 217, 205, 167,
			146, 59, 207, 190, 41, 63, 69, 184, 132, 75, 246, 217, 183, 15,
			230, 24, 28, 244, 120, 200, 65, 141, 70, 70, 182, 127, 96, 31,
			74, 224, 79, 16, 193, 225, 102, 79, 97, 129, 14, 162, 0, 215,
			235, 122, 145, 68, 80, 227, 115, 136, 37, 149, 128, 68, 149, 24,
			40, 97, 37, 32, 105, 37, 42, 182, 36, 146, 112, 201, 247, 60,
			20, 11, 3, 219, 117, 213, 246, 49, 58, 16, 9, 55, 110, 59,
			61, 133, 201, 18, 28, 135, 248, 47, 78, 41, 165, 24, 65, 97,
			33, 239, 10, 24, 14, 8, 38, 36, 67, 11, 94, 153, 151, 227,
			198, 165, 227, 209, 94, 177, 196, 72, 190, 24, 40, 206, 74, 210,
			220, 131, 57, 144, 195, 78, 199, 123, 130, 120, 148, 142, 255, 212,
			241, 29, 198, 1, 33, 209, 194, 220, 110, 123, 109, 174, 120, 235,
			72, 43, 83, 128, 241, 215, 67, 47, 228, 110, 25, 170, 160, 142,
			191, 84, 48, 72, 218, 147, 123, 223, 243, 16, 100, 79, 12, 125,
			55, 113, 229, 80, 114, 66, 147, 5, 91, 166, 210, 92, 216, 59,
			100, 168, 70, 17, 39, 32, 192, 93, 112, 160, 32, 205, 139, 161,
			132, 142, 180, 143, 136, 26, 216, 161, 28, 137, 217, 227, 12, 8,
			211, 33, 194, 113, 28, 62, 136, 96, 79, 68, 61, 146, 137, 99,
			213, 161, 65, 98, 131, 124, 65, 15, 132, 189, 162, 211, 145, 60,
			34, 184, 182, 33, 194, 228, 132, 179, 4, 115, 203, 139, 75, 215,
			177, 58, 44, 93, 107, 47, 46, 221, 92, 89, 188, 185, 116, 173,
			188, 184, 244, 96, 46, 142, 110, 9, 68, 167, 229, 101, 96, 203,
			136, 1, 245, 36, 249, 34, 24, 225, 230, 107, 37, 64, 110, 229,
			120, 1, 217, 251, 118, 203, 9, 189, 65, 84, 66, 180, 123, 4,
			170, 217, 128, 229, 49, 57, 119, 36, 148, 135, 208, 81, 5, 187,
			138, 71, 10, 127, 76, 87, 174, 29, 186, 12, 190, 137, 132, 213,
			106, 180, 104, 145, 21, 138, 199, 0, 212, 114, 95, 124, 239, 249,
			190, 77, 171, 139, 7, 11, 187, 173, 138, 43, 28, 89, 185, 199,
			247, 42, 35, 85, 42, 77, 222, 225, 33, 15, 28, 94, 217, 244,
			197, 158, 237, 63, 108, 144, 14, 178, 130, 10, 85, 198, 132, 20,
			89, 122, 132, 107, 37, 153, 166, 68, 235, 92, 169, 4, 143, 16,
			49, 162, 211, 203, 201, 195, 163, 196, 32, 52, 117, 143, 39, 214,
			114, 151, 29, 107, 34, 131, 111, 30, 201, 40, 236, 208, 208, 49,
			139, 132, 35, 203, 3, 149, 217, 208, 150, 229, 138, 239, 237, 133,
			118, 120, 72, 176, 187, 220, 139, 250, 254, 101, 122, 74, 198, 22,
			233, 204, 133, 165, 129, 156, 8, 145, 3, 238, 192, 149, 249, 251,
			11, 243, 253, 133, 121, 183, 61, 191, 117, 115, 254, 206, 205, 249,
			86, 121, 190, 243, 224, 74, 25, 182, 189, 199, 252, 192, 147, 156,
			182, 57, 232, 160, 209, 44, 13, 37, 87, 220, 110, 11, 215, 166,
			96, 189, 34, 225, 155, 71, 86, 171, 145, 128, 154, 13, 149, 172,
			220, 152, 44, 20, 31, 125, 91, 80, 39, 149, 113, 158, 251, 78,
			184, 106, 38, 240, 97, 129, 246, 11, 246, 192, 163, 9, 73, 90,
			213, 46, 66, 233, 90, 121, 145, 55, 217, 153, 8, 152, 95, 94,
			159, 95, 94, 103, 80, 68, 71, 138, 61, 58, 33, 180, 99, 59,
			35, 30, 130, 99, 15, 104, 129, 136, 142, 186, 42, 176, 213, 82,
			75, 150, 153, 84, 105, 57, 245, 63, 29, 114, 79, 171, 99, 238,
			236, 143, 218, 228, 57, 246, 231, 26, 203, 102, 51, 122, 198, 204,
			254, 95, 77, 159, 201, 255, 94, 131, 230, 104, 135, 155, 196, 190,
			232, 80, 200, 147, 143, 165, 23, 56, 227, 40, 139, 29, 15, 179,
			224, 206, 80, 70, 24, 11, 175, 218, 22, 177, 227, 246, 69, 15,
			192, 11, 28, 127, 40, 189, 125, 220, 40, 158, 98, 19, 168, 222,
			4, 233, 119, 34, 33, 53, 36, 39, 207, 36, 164, 129, 164, 121,
			158, 253, 179, 50, 70, 51, 179, 255, 95, 211, 205, 252, 159, 52,
			168, 139, 96, 33, 224, 93, 181, 15, 62, 178, 155, 182, 147, 93,
			35, 110, 36, 143, 223, 77, 215, 227, 129, 233, 6, 115, 223, 246,
			135, 92, 170, 35, 201, 17, 51, 58, 56, 149, 145, 231, 251, 208,
			179, 247, 57, 4, 227, 50, 137, 117, 60, 144, 169, 221, 155, 218,
			160, 119, 68, 136, 27, 227, 228, 244, 224, 121, 135, 197, 155, 198,
			82, 252, 203, 142, 113, 138, 54, 65, 118, 38, 78, 209, 200, 236,
			201, 83, 9, 105, 32, 121, 246, 92, 122, 147, 241, 199, 10, 187,
			222, 21, 101, 167, 23, 138, 190, 55, 236, 83, 144, 250, 67, 199,
			171, 116, 29, 138, 209, 74, 114, 89, 35, 43, 251, 75, 149, 248,
			166, 38, 190, 225, 152, 74, 223, 229, 95, 245, 61, 73, 254, 117,
			87, 37, 115, 239, 178, 236, 186, 39, 31, 155, 51, 108, 194, 235,
			219, 93, 78, 87, 34, 83, 77, 69, 204, 221, 101, 103, 235, 60,
			58, 16, 225, 99, 11, 183, 237, 29, 219, 225, 230, 251, 108, 154,
			246, 240, 129, 237, 63, 244, 6, 113, 127, 150, 52, 89, 3, 243,
			3, 118, 146, 63, 73, 59, 200, 89, 29, 140, 194, 84, 115, 58,
			105, 179, 6, 114, 238, 95, 13, 54, 105, 197, 22, 152, 167, 153,
			238, 185, 49, 31, 221, 115, 205, 60, 155, 236, 9, 25, 5, 118,
			159, 211, 69, 204, 84, 51, 165, 77, 147, 101, 177, 178, 204, 26,
			212, 78, 207, 230, 44, 59, 49, 8, 5, 166, 187, 217, 44, 53,
			39, 164, 121, 149, 157, 112, 66, 142, 181, 103, 118, 2, 180, 194,
			244, 114, 254, 249, 75, 156, 114, 90, 0, 155, 73, 87, 148, 239,
			123, 29, 202, 163, 179, 57, 186, 32, 74, 105, 124, 39, 15, 236,
			176, 239, 5, 221, 217, 19, 74, 183, 132, 54, 63, 101, 83, 142,
			8, 2, 238, 160, 188, 201, 215, 202, 27, 117, 70, 11, 104, 163,
			59, 140, 102, 167, 212, 141, 84, 76, 226, 27, 55, 180, 189, 128,
			187, 179, 12, 180, 194, 100, 51, 33, 205, 43, 236, 140, 35, 130,
			142, 215, 125, 24, 242, 125, 79, 122, 34, 152, 157, 38, 133, 78,
			171, 230, 102, 220, 106, 206, 179, 9, 215, 147, 143, 229, 236, 73,
			48, 10, 211, 203, 103, 202, 105, 232, 148, 113, 230, 155, 234, 173,
			121, 155, 153, 129, 154, 234, 135, 94, 50, 215, 114, 246, 20, 141,
			121, 103, 108, 204, 243, 241, 208, 60, 23, 60, 215, 34, 205, 139,
			44, 55, 8, 121, 199, 123, 50, 123, 154, 84, 138, 169, 185, 91,
			236, 212, 58, 247, 121, 196, 155, 234, 234, 242, 167, 76, 253, 220,
			167, 140, 109, 242, 232, 231, 140, 60, 100, 211, 219, 158, 76, 135,
			142, 180, 211, 198, 181, 51, 223, 99, 108, 96, 119, 249, 195, 72,
			60, 230, 65, 204, 100, 10, 91, 218, 216, 96, 190, 195, 136, 120,
			40, 189, 239, 85, 252, 77, 52, 39, 177, 161, 229, 125, 207, 145,
			103, 199, 243, 35, 30, 198, 33, 24, 83, 115, 255, 75, 99, 39,
			149, 108, 117, 55, 251, 82, 225, 75, 108, 180, 164, 105, 197, 76,
			47, 159, 31, 243, 122, 178, 88, 154, 163, 94, 230, 71, 236, 76,
			192, 159, 68, 15, 199, 148, 86, 203, 226, 20, 54, 239, 36, 138,
			47, 255, 81, 99, 83, 86, 58, 234, 38, 203, 169, 57, 48, 103,
			199, 35, 97, 124, 90, 242, 23, 95, 8, 91, 186, 172, 53, 151,
			152, 177, 201, 35, 243, 194, 216, 192, 209, 148, 228, 143, 211, 215,
			188, 206, 178, 104, 191, 121, 113, 236, 229, 216, 100, 228, 223, 122,
			161, 93, 57, 106, 245, 218, 131, 149, 55, 78, 142, 183, 82, 226,
			246, 159, 62, 100, 39, 204, 137, 108, 230, 71, 77, 99, 127, 173,
			209, 37, 105, 54, 99, 46, 255, 65, 59, 114, 223, 185, 116, 131,
			176, 249, 246, 238, 154, 5, 213, 97, 212, 19, 161, 44, 191, 228,
			210, 115, 87, 82, 49, 143, 175, 150, 70, 87, 132, 158, 132, 174,
			216, 231, 97, 128, 251, 150, 192, 141, 111, 188, 170, 3, 219, 65,
			198, 158, 195, 3, 68, 52, 119, 121, 136, 139, 16, 150, 203, 139,
			73, 177, 81, 136, 172, 35, 134, 129, 155, 156, 235, 110, 91, 107,
			181, 122, 171, 6, 29, 207, 231, 234, 52, 126, 210, 52, 114, 153,
			219, 201, 201, 58, 62, 170, 35, 250, 201, 140, 137, 141, 70, 198,
			52, 88, 102, 158, 30, 53, 211, 152, 206, 20, 217, 85, 58, 36,
			207, 158, 206, 152, 90, 190, 0, 85, 192, 85, 13, 182, 148, 194,
			241, 8, 124, 83, 141, 180, 97, 115, 173, 150, 126, 4, 48, 118,
			195, 126, 122, 242, 36, 59, 140, 129, 135, 113, 78, 63, 159, 247,
			201, 65, 84, 5, 94, 224, 66, 190, 64, 254, 99, 119, 194, 142,
			47, 134, 238, 248, 201, 172, 163, 54, 206, 10, 250, 134, 41, 214,
			13, 185, 140, 176, 156, 17, 103, 89, 241, 61, 137, 24, 232, 100,
			2, 42, 140, 115, 250, 100, 66, 105, 166, 113, 110, 234, 116, 66,
			25, 166, 113, 238, 156, 201, 170, 76, 207, 106, 102, 246, 66, 38,
			175, 229, 175, 65, 21, 226, 228, 3, 105, 206, 122, 51, 155, 53,
			205, 52, 46, 76, 206, 178, 58, 203, 102, 53, 180, 249, 45, 253,
			237, 124, 85, 217, 28, 87, 50, 176, 118, 192, 118, 93, 220, 154,
			30, 239, 129, 23, 68, 199, 134, 104, 100, 200, 91, 177, 33, 26,
			25, 242, 214, 212, 76, 66, 25, 166, 241, 214, 91, 179, 108, 135,
			36, 107, 166, 241, 182, 126, 57, 191, 70, 146, 147, 10, 57, 38,
			153, 255, 100, 217, 90, 22, 89, 166, 212, 132, 105, 188, 61, 125,
			46, 161, 80, 156, 121, 41, 161, 12, 211, 120, 251, 131, 57, 246,
			107, 166, 103, 117, 51, 123, 41, 179, 173, 229, 57, 84, 143, 184,
			11, 84, 49, 25, 134, 241, 85, 236, 19, 154, 175, 216, 79, 113,
			23, 219, 137, 134, 182, 239, 31, 170, 183, 18, 188, 78, 135, 162,
			58, 174, 167, 99, 167, 33, 184, 132, 120, 20, 79, 1, 218, 126,
			105, 242, 44, 251, 130, 101, 179, 58, 78, 1, 232, 103, 243, 203,
			196, 218, 90, 87, 75, 142, 195, 221, 59, 202, 222, 84, 216, 129,
			45, 83, 198, 8, 204, 98, 187, 117, 242, 57, 196, 62, 215, 201,
			231, 48, 53, 157, 80, 134, 105, 192, 233, 51, 236, 38, 137, 210,
			76, 227, 178, 126, 49, 191, 64, 162, 146, 34, 113, 188, 155, 199,
			162, 70, 113, 66, 127, 94, 78, 165, 160, 63, 47, 79, 157, 75,
			40, 195, 52, 46, 207, 92, 96, 79, 73, 138, 110, 26, 87, 116,
			51, 31, 144, 20, 218, 231, 191, 70, 194, 207, 94, 74, 180, 213,
			30, 95, 73, 186, 174, 79, 160, 244, 68, 77, 180, 248, 202, 212,
			169, 132, 50, 76, 227, 202, 217, 115, 236, 115, 82, 211, 48, 141,
			143, 245, 11, 249, 69, 82, 115, 115, 109, 7, 98, 236, 244, 134,
			254, 48, 38, 112, 124, 34, 8, 115, 209, 199, 83, 103, 19, 10,
			121, 159, 159, 97, 159, 145, 160, 172, 105, 44, 232, 133, 124, 133,
			4, 141, 66, 130, 142, 36, 94, 58, 197, 169, 156, 108, 14, 135,
			191, 147, 80, 154, 105, 44, 188, 123, 57, 161, 12, 211, 88, 248,
			232, 10, 115, 72, 206, 132, 105, 44, 233, 23, 242, 119, 73, 78,
			130, 219, 210, 12, 62, 230, 238, 170, 250, 96, 132, 7, 110, 252,
			21, 66, 218, 91, 157, 178, 165, 250, 96, 178, 163, 242, 56, 82,
			103, 130, 164, 156, 72, 40, 205, 52, 150, 38, 19, 179, 39, 12,
			211, 88, 58, 63, 195, 90, 164, 78, 206, 52, 174, 234, 23, 243,
			27, 71, 131, 45, 142, 238, 86, 12, 29, 233, 51, 44, 170, 32,
			227, 126, 136, 207, 92, 98, 152, 8, 145, 72, 197, 231, 38, 144,
			107, 226, 245, 156, 102, 26, 87, 211, 40, 204, 25, 166, 113, 117,
			230, 2, 219, 32, 241, 39, 76, 227, 186, 254, 113, 254, 198, 235,
			189, 158, 162, 81, 92, 231, 137, 98, 169, 196, 19, 57, 100, 148,
			248, 255, 132, 102, 26, 215, 223, 157, 79, 40, 195, 52, 174, 23,
			138, 236, 128, 36, 78, 154, 198, 45, 125, 38, 255, 93, 42, 81,
			12, 163, 227, 220, 111, 117, 32, 24, 9, 130, 61, 17, 65, 15,
			39, 62, 85, 99, 239, 48, 61, 71, 16, 195, 232, 245, 115, 50,
			57, 129, 146, 147, 57, 153, 212, 76, 227, 214, 228, 153, 132, 50,
			76, 227, 150, 121, 158, 125, 79, 42, 78, 153, 198, 23, 250, 76,
			190, 15, 247, 122, 60, 254, 86, 101, 220, 19, 200, 91, 129, 108,
			16, 33, 4, 2, 19, 94, 53, 109, 26, 133, 41, 238, 47, 247,
			56, 72, 187, 195, 253, 195, 68, 29, 8, 121, 215, 14, 93, 31,
			171, 134, 232, 164, 33, 149, 106, 57, 53, 129, 194, 115, 9, 165,
			153, 198, 23, 39, 78, 39, 148, 97, 26, 95, 156, 59, 207, 170,
			164, 37, 51, 141, 85, 253, 82, 254, 106, 124, 135, 140, 105, 24,
			18, 168, 255, 134, 171, 147, 77, 32, 143, 36, 78, 152, 102, 26,
			171, 83, 179, 9, 101, 152, 198, 234, 59, 239, 177, 235, 36, 108,
			218, 52, 106, 250, 59, 249, 143, 73, 24, 109, 8, 222, 80, 196,
			116, 22, 71, 166, 84, 206, 52, 106, 211, 137, 215, 167, 53, 211,
			168, 157, 189, 144, 80, 134, 105, 212, 102, 243, 108, 157, 4, 158,
			52, 141, 45, 253, 106, 254, 58, 9, 124, 161, 154, 189, 169, 244,
			147, 89, 100, 147, 82, 57, 211, 216, 154, 78, 130, 244, 164, 102,
			26, 91, 239, 150, 19, 202, 48, 141, 173, 165, 21, 246, 41, 73,
			63, 101, 26, 95, 234, 23, 242, 191, 34, 233, 10, 129, 191, 161,
			196, 83, 19, 56, 52, 113, 233, 41, 205, 52, 190, 156, 74, 236,
			61, 101, 152, 198, 151, 230, 12, 43, 51, 204, 73, 217, 70, 102,
			87, 203, 207, 65, 53, 249, 80, 19, 23, 150, 138, 18, 176, 131,
			231, 1, 9, 166, 206, 198, 228, 5, 182, 200, 178, 89, 3, 171,
			225, 87, 250, 217, 252, 229, 231, 170, 97, 26, 123, 41, 167, 88,
			47, 131, 202, 223, 87, 177, 94, 6, 149, 191, 175, 226, 242, 103,
			80, 249, 251, 234, 244, 25, 246, 9, 241, 214, 76, 163, 173, 95,
			204, 23, 143, 205, 72, 175, 144, 128, 165, 175, 157, 74, 192, 210,
			215, 142, 147, 142, 65, 165, 175, 61, 115, 129, 64, 104, 214, 204,
			126, 157, 249, 159, 10, 132, 142, 89, 222, 229, 17, 125, 244, 128,
			24, 65, 93, 79, 28, 177, 31, 83, 250, 215, 147, 38, 91, 96,
			217, 108, 22, 237, 127, 160, 159, 205, 195, 43, 236, 239, 242, 164,
			220, 101, 201, 248, 7, 177, 106, 89, 50, 254, 65, 108, 124, 150,
			140, 127, 112, 250, 12, 91, 33, 198, 154, 105, 124, 171, 95, 204,
			127, 244, 90, 227, 199, 217, 163, 229, 223, 166, 236, 209, 242, 111,
			99, 203, 179, 100, 249, 183, 51, 23, 216, 10, 211, 179, 19, 102,
			214, 206, 12, 180, 252, 149, 163, 150, 211, 103, 121, 47, 216, 45,
			99, 195, 177, 120, 216, 147, 231, 105, 226, 39, 208, 112, 71, 159,
			137, 39, 62, 14, 204, 132, 69, 58, 18, 58, 34, 140, 149, 155,
			32, 219, 157, 88, 185, 9, 178, 221, 137, 3, 114, 130, 108, 119,
			204, 243, 236, 255, 105, 196, 92, 51, 141, 174, 62, 155, 255, 173,
			70, 220, 233, 8, 12, 77, 127, 110, 19, 9, 33, 119, 56, 125,
			84, 71, 71, 155, 227, 155, 178, 50, 238, 134, 220, 100, 54, 233,
			92, 28, 7, 3, 14, 166, 34, 154, 218, 134, 233, 157, 206, 161,
			74, 216, 85, 93, 120, 116, 188, 80, 170, 190, 169, 242, 232, 217,
			110, 170, 60, 122, 182, 59, 117, 62, 161, 12, 211, 232, 94, 124,
			139, 109, 145, 238, 186, 105, 124, 167, 95, 204, 223, 34, 213, 251,
			246, 19, 175, 63, 236, 67, 48, 236, 239, 241, 16, 37, 135, 92,
			14, 253, 136, 62, 106, 161, 115, 57, 151, 39, 123, 171, 228, 171,
			232, 84, 38, 34, 132, 239, 226, 58, 49, 65, 216, 232, 187, 201,
			115, 9, 101, 152, 198, 119, 51, 23, 216, 95, 41, 135, 25, 166,
			33, 244, 153, 252, 95, 104, 80, 5, 181, 161, 167, 139, 143, 193,
			192, 63, 84, 37, 212, 127, 97, 82, 33, 190, 209, 243, 85, 245,
			18, 129, 127, 8, 114, 56, 24, 8, 117, 250, 30, 115, 241, 36,
			204, 81, 150, 45, 211, 62, 232, 243, 207, 232, 207, 255, 152, 139,
			175, 200, 98, 146, 190, 250, 30, 251, 140, 50, 74, 119, 100, 145,
			72, 89, 141, 199, 2, 162, 49, 145, 186, 19, 83, 138, 72, 99,
			1, 209, 152, 48, 207, 179, 79, 25, 130, 169, 172, 204, 60, 211,
			242, 37, 10, 212, 228, 171, 113, 1, 246, 11, 97, 251, 124, 180,
			34, 214, 144, 147, 51, 20, 173, 57, 140, 214, 225, 11, 209, 58,
			182, 144, 212, 119, 66, 35, 13, 115, 20, 173, 195, 88, 195, 28,
			69, 235, 48, 214, 48, 71, 209, 58, 52, 207, 179, 37, 226, 173,
			153, 198, 19, 125, 46, 255, 225, 145, 189, 134, 132, 190, 29, 57,
			189, 228, 246, 74, 137, 76, 153, 227, 214, 231, 137, 158, 82, 57,
			211, 120, 50, 125, 62, 161, 144, 223, 204, 123, 9, 101, 152, 198,
			19, 248, 128, 253, 94, 35, 89, 186, 105, 252, 160, 191, 151, 255,
			205, 248, 194, 136, 63, 183, 143, 191, 181, 31, 91, 30, 99, 171,
			226, 72, 114, 163, 153, 26, 91, 13, 236, 248, 229, 16, 209, 12,
			163, 95, 2, 1, 125, 17, 242, 35, 62, 86, 234, 97, 136, 254,
			144, 122, 9, 93, 241, 67, 92, 183, 115, 20, 162, 63, 188, 243,
			46, 187, 193, 244, 92, 198, 204, 253, 70, 203, 252, 168, 105, 249,
			95, 65, 53, 253, 150, 191, 35, 66, 232, 219, 129, 55, 24, 250,
			234, 246, 97, 124, 67, 135, 66, 166, 153, 145, 203, 104, 102, 246,
			55, 218, 228, 57, 220, 19, 229, 232, 186, 225, 183, 154, 254, 89,
			190, 4, 234, 36, 40, 78, 254, 114, 188, 82, 129, 45, 15, 3,
			167, 23, 138, 64, 12, 165, 127, 168, 142, 189, 115, 234, 240, 255,
			183, 90, 238, 100, 66, 234, 72, 158, 202, 39, 164, 129, 228, 252,
			77, 156, 214, 28, 93, 5, 252, 111, 77, 47, 230, 47, 195, 38,
			143, 210, 255, 74, 240, 146, 194, 160, 56, 104, 52, 38, 55, 149,
			144, 58, 146, 108, 38, 33, 13, 36, 223, 191, 130, 33, 153, 203,
			232, 186, 153, 253, 157, 166, 87, 242, 115, 52, 69, 169, 128, 99,
			211, 175, 98, 128, 42, 253, 78, 203, 165, 36, 113, 152, 190, 152,
			144, 6, 146, 31, 44, 36, 167, 246, 255, 25, 0, 0, 255, 255,
			84, 241, 117, 39, 194, 57, 0, 0},
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
	ret, err := discovery.GetDescriptorSet("instances.Instances")
	if err != nil {
		panic(err)
	}
	return ret
}
