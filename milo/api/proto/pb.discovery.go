// Code generated by cproto. DO NOT EDIT.

package milo

import discovery "go.chromium.org/luci/grpc/discovery"

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"milo.Buildbot",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 188, 122, 75, 119, 19, 73,
			150, 191, 242, 97, 35, 7, 239, 196, 174, 2, 153, 130, 139, 10,
			202, 82, 151, 156, 178, 77, 65, 129, 169, 199, 145, 229, 196, 36,
			45, 75, 234, 148, 92, 52, 212, 169, 3, 105, 101, 200, 202, 238,
			84, 134, 58, 35, 101, 99, 248, 211, 181, 234, 255, 114, 86, 179,
			156, 213, 44, 102, 209, 155, 217, 76, 127, 128, 153, 245, 172, 103,
			61, 155, 249, 14, 211, 187, 57, 55, 34, 82, 146, 121, 116, 117,
			247, 204, 105, 29, 56, 228, 141, 140, 184, 247, 119, 31, 113, 239,
			141, 72, 200, 63, 87, 201, 245, 3, 198, 14, 34, 90, 29, 37,
			44, 101, 251, 227, 126, 53, 13, 135, 148, 167, 254, 112, 100, 139,
			33, 235, 188, 156, 96, 103, 19, 138, 15, 200, 66, 55, 155, 99,
			93, 38, 167, 56, 237, 177, 56, 224, 151, 53, 208, 74, 134, 151,
			145, 214, 34, 153, 139, 253, 152, 241, 203, 58, 104, 165, 57, 79,
			18, 91, 191, 37, 151, 122, 108, 104, 191, 197, 115, 235, 220, 132,
			99, 27, 135, 218, 218, 179, 207, 15, 194, 116, 48, 222, 183, 123,
			108, 88, 61, 96, 145, 31, 31, 76, 33, 142, 210, 227, 17, 229,
			83, 164, 255, 173, 105, 255, 168, 27, 59, 237, 173, 223, 235, 215,
			118, 36, 231, 182, 154, 107, 63, 161, 81, 244, 243, 152, 29, 197,
			93, 92, 243, 248, 15, 171, 228, 148, 53, 119, 45, 247, 255, 53,
			141, 252, 251, 25, 162, 157, 177, 140, 107, 57, 107, 227, 95, 207,
			128, 88, 209, 99, 17, 108, 141, 251, 125, 154, 112, 88, 5, 201,
			107, 133, 67, 224, 167, 62, 132, 113, 74, 147, 222, 192, 143, 15,
			40, 244, 89, 50, 244, 83, 2, 117, 54, 58, 78, 194, 131, 65,
			10, 27, 107, 107, 247, 212, 2, 112, 227, 158, 13, 80, 139, 34,
			16, 239, 56, 36, 148, 211, 228, 144, 6, 54, 129, 65, 154, 142,
			248, 102, 181, 26, 208, 67, 26, 177, 17, 77, 120, 102, 12, 212,
			116, 164, 64, 172, 238, 75, 16, 85, 66, 192, 163, 65, 200, 211,
			36, 220, 31, 167, 33, 139, 193, 143, 3, 24, 115, 10, 97, 12,
			156, 141, 147, 30, 21, 35, 251, 97, 236, 39, 199, 2, 23, 175,
			192, 81, 152, 14, 128, 37, 226, 95, 54, 78, 9, 12, 89, 16,
			246, 195, 158, 143, 28, 42, 224, 39, 20, 70, 52, 25, 134, 105,
			74, 3, 24, 37, 236, 48, 12, 104, 0, 233, 192, 79, 33, 29,
			160, 118, 81, 196, 142, 194, 248, 0, 208, 149, 33, 46, 226, 184,
			136, 192, 144, 166, 155, 132, 0, 254, 126, 246, 22, 48, 14, 172,
			159, 33, 234, 177, 128, 194, 112, 204, 83, 72, 104, 234, 135, 177,
			224, 234, 239, 179, 67, 124, 165, 44, 70, 32, 102, 105, 216, 163,
			21, 72, 7, 33, 135, 40, 228, 41, 114, 152, 149, 24, 7, 111,
			193, 9, 66, 222, 139, 252, 112, 72, 19, 251, 67, 32, 194, 120,
			214, 22, 25, 136, 81, 194, 130, 113, 143, 78, 113, 144, 41, 144,
			255, 21, 14, 2, 74, 187, 128, 245, 198, 67, 26, 167, 126, 230,
			164, 42, 75, 128, 165, 3, 154, 192, 208, 79, 105, 18, 250, 17,
			159, 154, 90, 56, 40, 29, 80, 2, 179, 232, 39, 74, 53, 105,
			40, 86, 34, 227, 216, 31, 82, 4, 52, 27, 91, 49, 155, 190,
			19, 118, 15, 83, 142, 26, 197, 146, 21, 75, 56, 12, 253, 99,
			216, 167, 24, 41, 1, 164, 12, 104, 28, 176, 132, 83, 12, 138,
			81, 194, 134, 44, 165, 32, 109, 146, 114, 8, 104, 18, 30, 210,
			0, 250, 9, 27, 18, 105, 5, 206, 250, 233, 17, 134, 137, 138,
			32, 224, 35, 218, 195, 8, 130, 81, 18, 98, 96, 37, 24, 59,
			177, 140, 34, 206, 5, 118, 2, 221, 71, 110, 7, 58, 173, 135,
			221, 39, 53, 207, 1, 183, 3, 109, 175, 245, 157, 187, 237, 108,
			195, 214, 83, 232, 62, 114, 160, 222, 106, 63, 245, 220, 157, 71,
			93, 120, 212, 106, 108, 59, 94, 7, 106, 205, 109, 168, 183, 154,
			93, 207, 221, 218, 235, 182, 188, 14, 129, 98, 173, 3, 110, 167,
			40, 222, 212, 154, 79, 193, 249, 101, 219, 115, 58, 29, 104, 121,
			224, 238, 182, 27, 174, 179, 13, 79, 106, 158, 87, 107, 118, 93,
			167, 83, 1, 183, 89, 111, 236, 109, 187, 205, 157, 10, 108, 237,
			117, 161, 217, 234, 18, 104, 184, 187, 110, 215, 217, 134, 110, 171,
			34, 196, 190, 187, 14, 90, 15, 97, 215, 241, 234, 143, 106, 205,
			110, 109, 203, 109, 184, 221, 167, 66, 224, 67, 183, 219, 68, 97,
			15, 91, 30, 129, 26, 180, 107, 94, 215, 173, 239, 53, 106, 30,
			180, 247, 188, 118, 171, 227, 0, 106, 182, 237, 118, 234, 141, 154,
			187, 235, 108, 219, 224, 54, 161, 217, 2, 231, 59, 167, 217, 133,
			206, 163, 90, 163, 113, 82, 81, 2, 173, 39, 77, 199, 67, 244,
			179, 106, 194, 150, 3, 13, 183, 182, 213, 112, 80, 148, 208, 115,
			219, 245, 156, 122, 23, 21, 154, 62, 213, 221, 109, 167, 217, 173,
			53, 42, 4, 58, 109, 167, 238, 214, 26, 21, 112, 126, 233, 236,
			182, 27, 53, 239, 105, 69, 49, 237, 56, 191, 216, 115, 154, 93,
			183, 214, 128, 237, 218, 110, 109, 199, 233, 64, 233, 167, 172, 210,
			246, 90, 245, 61, 207, 217, 69, 212, 173, 135, 208, 217, 219, 234,
			116, 221, 238, 94, 215, 129, 157, 86, 107, 91, 24, 187, 227, 120,
			223, 185, 117, 167, 243, 0, 26, 173, 142, 48, 216, 94, 199, 169,
			16, 216, 174, 117, 107, 66, 116, 219, 107, 61, 116, 187, 157, 7,
			248, 188, 181, 215, 113, 133, 225, 220, 102, 215, 241, 188, 189, 118,
			215, 109, 53, 203, 240, 168, 245, 196, 249, 206, 241, 160, 94, 219,
			235, 56, 219, 194, 194, 173, 38, 106, 139, 177, 226, 180, 188, 167,
			200, 22, 237, 32, 60, 80, 129, 39, 143, 156, 238, 35, 199, 67,
			163, 10, 107, 213, 208, 12, 157, 174, 231, 214, 187, 179, 211, 90,
			30, 116, 91, 94, 151, 204, 232, 9, 77, 103, 167, 225, 238, 56,
			205, 186, 131, 175, 91, 200, 230, 137, 219, 113, 202, 80, 243, 220,
			14, 78, 112, 133, 96, 120, 82, 123, 10, 173, 61, 161, 53, 58,
			106, 175, 227, 16, 249, 60, 19, 186, 21, 225, 79, 112, 31, 66,
			109, 251, 59, 23, 145, 171, 217, 237, 86, 167, 227, 170, 112, 17,
			102, 171, 63, 82, 54, 183, 9, 201, 19, 77, 183, 12, 200, 127,
			140, 79, 121, 203, 40, 230, 30, 144, 5, 162, 231, 111, 201, 71,
			57, 248, 105, 238, 186, 24, 188, 46, 31, 229, 224, 205, 220, 150,
			24, 60, 45, 31, 229, 224, 173, 92, 69, 12, 106, 242, 81, 14,
			126, 150, 171, 138, 65, 245, 40, 7, 87, 114, 69, 49, 72, 228,
			163, 28, 44, 229, 110, 136, 193, 155, 242, 241, 95, 150, 137, 110,
			230, 172, 185, 87, 88, 249, 10, 255, 180, 12, 53, 152, 148, 92,
			145, 31, 41, 167, 113, 202, 193, 135, 17, 11, 227, 84, 100, 181,
			112, 136, 85, 38, 160, 35, 26, 7, 52, 22, 89, 209, 143, 143,
			229, 248, 43, 22, 83, 130, 217, 164, 231, 71, 52, 14, 252, 164,
			50, 229, 66, 3, 240, 57, 168, 62, 64, 100, 207, 126, 226, 247,
			166, 53, 34, 123, 129, 37, 0, 155, 2, 65, 99, 141, 100, 145,
			44, 113, 97, 12, 123, 221, 58, 56, 35, 214, 27, 8, 113, 54,
			184, 41, 132, 28, 104, 140, 149, 5, 235, 31, 102, 97, 145, 63,
			219, 9, 139, 232, 40, 13, 123, 176, 147, 208, 3, 150, 132, 126,
			12, 117, 133, 9, 142, 6, 97, 111, 0, 244, 101, 74, 81, 32,
			102, 204, 233, 164, 12, 56, 129, 125, 191, 247, 235, 35, 63, 193,
			25, 12, 142, 169, 159, 0, 139, 223, 17, 233, 115, 62, 30, 162,
			84, 63, 138, 96, 24, 198, 227, 148, 138, 154, 8, 119, 215, 200,
			68, 165, 136, 197, 7, 21, 8, 109, 106, 67, 68, 253, 209, 84,
			213, 132, 66, 145, 15, 169, 159, 208, 160, 8, 156, 201, 82, 27,
			179, 217, 89, 4, 82, 127, 63, 162, 40, 51, 166, 20, 69, 246,
			89, 34, 155, 142, 17, 86, 81, 81, 32, 192, 19, 237, 71, 200,
			85, 178, 94, 91, 91, 91, 95, 21, 127, 186, 107, 107, 155, 226,
			207, 51, 212, 226, 254, 253, 251, 247, 87, 215, 55, 86, 111, 175,
			119, 55, 110, 111, 222, 185, 191, 121, 231, 190, 125, 63, 251, 61,
			179, 9, 108, 29, 163, 193, 211, 36, 236, 165, 194, 148, 10, 82,
			130, 236, 43, 112, 68, 129, 198, 124, 156, 80, 57, 122, 68, 161,
			135, 22, 99, 241, 33, 77, 82, 72, 25, 81, 94, 101, 67, 0,
			239, 97, 29, 110, 223, 190, 125, 31, 155, 36, 10, 200, 50, 62,
			224, 54, 129, 14, 165, 240, 125, 214, 237, 28, 29, 29, 217, 33,
			77, 251, 54, 75, 14, 170, 73, 191, 135, 127, 113, 145, 157, 190,
			76, 127, 40, 253, 57, 179, 202, 88, 96, 62, 5, 231, 165, 63,
			28, 69, 148, 19, 146, 61, 194, 250, 38, 212, 217, 112, 52, 78,
			233, 76, 72, 11, 108, 237, 86, 199, 253, 37, 188, 192, 8, 42,
			149, 95, 216, 170, 113, 153, 78, 154, 52, 144, 15, 228, 155, 105,
			235, 203, 105, 250, 92, 57, 175, 36, 150, 55, 247, 26, 141, 114,
			249, 189, 243, 68, 12, 151, 214, 202, 15, 102, 48, 109, 252, 20,
			166, 3, 154, 34, 23, 214, 15, 252, 227, 25, 108, 60, 77, 198,
			189, 84, 8, 56, 244, 35, 72, 15, 149, 196, 19, 211, 63, 75,
			15, 43, 32, 0, 61, 248, 107, 85, 58, 180, 211, 67, 164, 254,
			148, 70, 114, 210, 152, 211, 30, 252, 12, 214, 215, 214, 78, 106,
			120, 251, 131, 26, 62, 9, 227, 219, 27, 240, 98, 135, 166, 157,
			99, 158, 210, 33, 190, 174, 241, 135, 97, 68, 187, 39, 29, 241,
			208, 109, 56, 93, 119, 215, 129, 126, 170, 96, 124, 104, 205, 103,
			253, 52, 67, 186, 231, 54, 187, 119, 191, 128, 52, 236, 253, 154,
			195, 215, 80, 42, 149, 228, 72, 185, 159, 218, 193, 209, 163, 240,
			96, 176, 237, 167, 98, 85, 25, 190, 250, 10, 110, 111, 148, 225,
			255, 129, 120, 215, 96, 71, 217, 171, 204, 110, 213, 42, 212, 16,
			111, 192, 142, 184, 96, 137, 59, 107, 125, 109, 109, 38, 47, 113,
			123, 50, 129, 138, 124, 180, 126, 247, 221, 45, 55, 225, 134, 203,
			215, 239, 126, 241, 197, 23, 95, 222, 190, 187, 182, 54, 217, 255,
			251, 180, 207, 18, 10, 123, 113, 248, 50, 227, 114, 255, 203, 181,
			183, 185, 216, 127, 157, 51, 75, 82, 127, 40, 149, 164, 81, 170,
			194, 89, 248, 43, 195, 234, 44, 156, 159, 136, 96, 228, 131, 230,
			202, 248, 220, 154, 225, 35, 2, 160, 124, 34, 0, 190, 248, 96,
			0, 60, 246, 15, 125, 120, 33, 29, 105, 247, 198, 73, 66, 227,
			20, 167, 236, 134, 81, 20, 242, 153, 0, 192, 116, 9, 67, 49,
			10, 95, 195, 135, 23, 252, 137, 48, 135, 175, 167, 163, 118, 76,
			143, 182, 198, 97, 20, 208, 164, 84, 70, 197, 58, 202, 66, 74,
			132, 52, 76, 89, 242, 194, 31, 206, 105, 74, 221, 195, 56, 69,
			205, 213, 76, 169, 186, 82, 91, 88, 160, 108, 239, 35, 103, 129,
			101, 106, 131, 59, 31, 180, 129, 210, 34, 43, 162, 208, 62, 78,
			7, 178, 73, 62, 97, 254, 89, 248, 165, 242, 219, 190, 217, 161,
			105, 125, 106, 141, 82, 89, 100, 192, 199, 157, 86, 19, 118, 253,
			209, 40, 140, 15, 8, 1, 55, 150, 35, 242, 68, 90, 17, 69,
			110, 198, 78, 199, 35, 122, 178, 138, 129, 175, 114, 180, 58, 184,
			16, 248, 62, 203, 224, 127, 102, 34, 86, 162, 108, 232, 98, 109,
			8, 121, 69, 178, 145, 163, 40, 172, 248, 26, 139, 232, 155, 213,
			215, 67, 22, 167, 131, 55, 171, 175, 3, 255, 248, 77, 247, 245,
			128, 141, 147, 55, 155, 175, 135, 97, 252, 102, 243, 53, 167, 189,
			55, 223, 219, 175, 177, 49, 192, 64, 126, 243, 195, 179, 34, 129,
			163, 1, 77, 40, 200, 213, 200, 200, 143, 142, 252, 99, 14, 244,
			37, 54, 22, 124, 82, 247, 251, 108, 156, 64, 16, 30, 132, 41,
			199, 10, 31, 81, 80, 146, 42, 32, 68, 85, 8, 72, 97, 21,
			16, 210, 42, 162, 90, 9, 145, 162, 18, 191, 162, 9, 91, 29,
			249, 65, 32, 143, 70, 233, 17, 203, 184, 81, 191, 55, 64, 189,
			232, 164, 99, 241, 163, 73, 117, 175, 168, 118, 2, 75, 225, 1,
			131, 241, 72, 20, 218, 108, 105, 73, 84, 125, 57, 184, 254, 254,
			190, 166, 92, 33, 66, 62, 27, 73, 206, 82, 82, 241, 89, 17,
			248, 184, 223, 15, 95, 98, 179, 133, 103, 116, 42, 91, 21, 140,
			3, 108, 179, 160, 84, 220, 235, 214, 139, 229, 7, 39, 70, 9,
			26, 40, 161, 191, 25, 135, 9, 13, 108, 168, 129, 184, 58, 184,
			45, 131, 129, 139, 243, 102, 248, 138, 38, 192, 7, 108, 28, 5,
			153, 41, 199, 156, 138, 214, 170, 228, 243, 137, 180, 0, 246, 143,
			9, 194, 40, 163, 3, 98, 60, 225, 197, 169, 234, 175, 222, 14,
			37, 52, 164, 127, 66, 212, 200, 79, 248, 84, 204, 62, 37, 32,
			186, 152, 148, 129, 223, 235, 209, 81, 10, 251, 44, 29, 8, 153,
			184, 86, 30, 136, 51, 29, 248, 59, 56, 192, 143, 129, 245, 251,
			156, 202, 122, 255, 144, 37, 64, 229, 94, 171, 64, 113, 99, 109,
			253, 75, 204, 153, 235, 119, 186, 107, 235, 155, 183, 215, 54, 215,
			239, 216, 107, 235, 207, 138, 42, 186, 57, 8, 122, 146, 116, 71,
			62, 79, 9, 136, 153, 66, 62, 139, 225, 177, 31, 143, 253, 228,
			24, 214, 239, 84, 0, 185, 217, 106, 3, 249, 135, 126, 167, 151,
			132, 163, 180, 130, 173, 223, 137, 102, 199, 7, 44, 26, 192, 246,
			127, 69, 177, 48, 51, 121, 62, 86, 193, 62, 211, 135, 242, 212,
			199, 110, 50, 128, 239, 83, 230, 118, 90, 29, 177, 199, 74, 229,
			233, 158, 154, 92, 248, 216, 67, 246, 42, 140, 34, 95, 108, 46,
			26, 175, 238, 117, 170, 1, 235, 241, 234, 19, 186, 95, 157, 34,
			169, 122, 180, 79, 19, 26, 247, 104, 117, 39, 98, 251, 126, 244,
			188, 37, 32, 240, 42, 226, 169, 206, 8, 249, 65, 92, 203, 12,
			88, 96, 163, 46, 50, 209, 84, 196, 54, 87, 136, 94, 96, 103,
			38, 218, 232, 236, 225, 69, 166, 15, 106, 186, 79, 51, 101, 41,
			54, 161, 239, 211, 240, 251, 23, 60, 77, 250, 98, 229, 140, 66,
			172, 199, 237, 145, 204, 107, 168, 202, 70, 53, 10, 247, 19, 63,
			57, 22, 23, 115, 246, 32, 29, 70, 159, 138, 167, 108, 109, 153,
			76, 238, 61, 100, 94, 84, 50, 248, 136, 246, 96, 229, 214, 211,
			213, 91, 195, 213, 91, 65, 247, 214, 163, 205, 91, 187, 155, 183,
			58, 246, 173, 254, 179, 21, 27, 26, 225, 175, 233, 81, 200, 105,
			5, 19, 22, 218, 71, 248, 136, 8, 232, 24, 206, 200, 237, 49,
			11, 124, 17, 170, 43, 28, 190, 127, 225, 118, 90, 89, 161, 127,
			40, 83, 85, 160, 200, 82, 249, 197, 15, 37, 121, 7, 167, 178,
			220, 175, 88, 32, 29, 129, 15, 171, 136, 170, 234, 143, 66, 225,
			143, 108, 84, 168, 83, 149, 88, 171, 239, 242, 22, 122, 102, 2,
			86, 87, 9, 148, 209, 134, 108, 95, 220, 123, 249, 74, 199, 148,
			226, 73, 105, 36, 182, 6, 235, 195, 1, 141, 105, 226, 203, 77,
			150, 109, 48, 46, 19, 242, 196, 244, 54, 17, 63, 195, 204, 105,
			150, 241, 42, 127, 145, 252, 131, 70, 76, 51, 167, 231, 44, 227,
			71, 125, 177, 240, 119, 26, 120, 211, 99, 91, 22, 243, 172, 47,
			66, 93, 88, 151, 135, 113, 111, 182, 231, 32, 239, 111, 58, 96,
			119, 204, 83, 12, 2, 81, 183, 62, 112, 160, 32, 239, 59, 81,
			60, 131, 48, 238, 69, 99, 30, 30, 82, 155, 144, 179, 100, 14,
			209, 153, 150, 249, 163, 254, 234, 18, 57, 35, 201, 57, 68, 123,
			42, 163, 52, 203, 248, 49, 127, 62, 163, 12, 203, 248, 209, 186,
			68, 254, 75, 234, 165, 89, 230, 239, 52, 221, 42, 252, 135, 6,
			77, 22, 175, 198, 244, 192, 79, 195, 67, 122, 242, 236, 232, 43,
			77, 1, 143, 79, 239, 203, 177, 54, 52, 213, 194, 44, 111, 195,
			161, 31, 141, 41, 151, 161, 55, 101, 38, 46, 6, 121, 26, 70,
			17, 12, 252, 67, 10, 241, 172, 76, 193, 90, 45, 36, 242, 12,
			212, 99, 227, 56, 69, 215, 224, 73, 49, 59, 30, 191, 109, 60,
			117, 244, 170, 168, 191, 228, 132, 129, 206, 9, 173, 53, 211, 154,
			251, 157, 166, 255, 184, 168, 12, 166, 205, 9, 189, 79, 101, 164,
			48, 67, 254, 108, 70, 26, 72, 94, 184, 184, 63, 47, 115, 46,
			249, 55, 32, 27, 7, 204, 238, 13, 18, 54, 12, 199, 67, 17,
			186, 209, 184, 23, 86, 135, 97, 196, 48, 116, 229, 21, 114, 85,
			116, 43, 251, 44, 85, 151, 248, 38, 190, 46, 252, 212, 93, 127,
			113, 76, 206, 238, 250, 60, 165, 137, 71, 127, 51, 166, 60, 181,
			44, 98, 198, 254, 144, 138, 187, 253, 5, 79, 60, 91, 171, 196,
			162, 47, 123, 209, 56, 160, 207, 3, 140, 64, 145, 187, 47, 19,
			208, 74, 121, 239, 162, 122, 179, 61, 121, 97, 221, 32, 103, 98,
			246, 156, 14, 199, 145, 56, 194, 94, 62, 45, 38, 158, 142, 153,
			147, 13, 21, 127, 75, 22, 177, 141, 146, 69, 94, 2, 192, 218,
			98, 21, 72, 94, 156, 127, 99, 63, 18, 8, 242, 222, 132, 182,
			238, 146, 188, 188, 212, 166, 129, 248, 194, 112, 122, 163, 240, 246,
			87, 5, 123, 82, 188, 188, 201, 92, 212, 40, 240, 83, 255, 178,
			1, 90, 233, 140, 39, 158, 139, 127, 175, 145, 197, 45, 101, 48,
			241, 111, 166, 254, 71, 100, 126, 40, 224, 40, 3, 40, 202, 186,
			76, 78, 237, 203, 70, 83, 200, 94, 240, 50, 210, 90, 38, 11,
			226, 241, 121, 60, 30, 10, 25, 134, 151, 23, 3, 205, 241, 240,
			47, 180, 92, 113, 133, 92, 60, 129, 74, 216, 36, 195, 175, 205,
			224, 255, 79, 141, 44, 157, 152, 201, 255, 122, 5, 22, 201, 92,
			20, 14, 195, 84, 128, 159, 243, 36, 97, 173, 144, 243, 34, 148,
			3, 250, 92, 117, 182, 151, 77, 1, 251, 156, 26, 86, 125, 234,
			255, 125, 112, 60, 54, 243, 115, 23, 230, 139, 63, 39, 214, 73,
			13, 133, 49, 170, 100, 94, 224, 230, 151, 53, 48, 74, 167, 55,
			62, 182, 49, 206, 237, 119, 172, 230, 169, 105, 143, 205, 188, 126,
			193, 216, 248, 163, 70, 242, 217, 28, 171, 65, 62, 198, 62, 251,
			125, 241, 119, 73, 178, 59, 177, 37, 10, 5, 57, 248, 190, 5,
			197, 156, 181, 75, 22, 119, 104, 250, 174, 219, 10, 239, 65, 150,
			113, 252, 16, 234, 98, 206, 106, 147, 165, 183, 217, 73, 205, 151,
			223, 179, 38, 243, 122, 225, 242, 251, 94, 74, 142, 143, 255, 248,
			49, 153, 183, 76, 51, 247, 11, 141, 252, 65, 19, 223, 191, 204,
			156, 181, 241, 123, 237, 196, 167, 172, 245, 187, 162, 53, 109, 236,
			213, 93, 168, 141, 211, 1, 75, 184, 253, 129, 239, 89, 123, 92,
			84, 52, 245, 213, 96, 250, 245, 39, 228, 112, 192, 14, 105, 18,
			99, 219, 30, 7, 234, 99, 70, 109, 228, 247, 144, 113, 216, 163,
			49, 150, 244, 239, 104, 194, 67, 22, 195, 134, 189, 150, 229, 89,
			217, 145, 244, 217, 56, 14, 178, 111, 43, 13, 183, 238, 52, 59,
			14, 244, 195, 136, 78, 46, 90, 231, 243, 103, 200, 2, 209, 141,
			156, 101, 228, 79, 149, 200, 83, 162, 207, 231, 44, 243, 108, 238,
			162, 86, 216, 133, 109, 167, 237, 57, 245, 90, 215, 217, 182, 97,
			155, 65, 204, 48, 215, 251, 201, 180, 99, 11, 57, 212, 218, 174,
			248, 104, 49, 160, 144, 101, 75, 108, 156, 15, 195, 30, 133, 128,
			246, 195, 56, 148, 223, 100, 176, 8, 207, 99, 233, 58, 155, 191,
			64, 78, 19, 115, 94, 212, 224, 115, 122, 19, 235, 216, 188, 172,
			106, 231, 230, 63, 201, 40, 221, 50, 206, 93, 251, 60, 163, 12,
			203, 56, 119, 247, 177, 90, 166, 89, 198, 121, 189, 165, 94, 105,
			72, 205, 23, 50, 74, 183, 140, 243, 203, 213, 140, 50, 44, 227,
			252, 230, 207, 213, 50, 221, 50, 46, 232, 158, 122, 133, 76, 46,
			204, 47, 103, 20, 190, 187, 186, 158, 81, 134, 101, 92, 248, 170,
			73, 190, 18, 119, 189, 230, 98, 238, 170, 86, 88, 19, 10, 38,
			50, 50, 196, 39, 34, 63, 140, 179, 238, 62, 251, 182, 132, 207,
			50, 51, 216, 211, 158, 99, 49, 191, 132, 0, 100, 203, 177, 164,
			91, 51, 21, 126, 73, 95, 252, 120, 166, 194, 47, 233, 249, 153,
			10, 191, 180, 112, 118, 166, 194, 47, 93, 184, 72, 14, 85, 129,
			55, 174, 232, 215, 11, 33, 132, 125, 72, 147, 49, 173, 128, 202,
			7, 24, 78, 35, 22, 115, 42, 191, 179, 206, 124, 134, 28, 199,
			129, 252, 158, 150, 82, 127, 40, 78, 132, 61, 172, 207, 209, 49,
			164, 201, 49, 158, 127, 177, 224, 78, 242, 204, 38, 129, 85, 224,
			17, 86, 243, 48, 238, 179, 172, 37, 209, 76, 203, 188, 162, 47,
			89, 10, 148, 54, 135, 56, 230, 51, 10, 81, 157, 186, 146, 81,
			134, 101, 92, 185, 122, 141, 172, 9, 192, 186, 101, 44, 235, 87,
			10, 159, 130, 155, 1, 78, 199, 137, 56, 159, 192, 36, 99, 193,
			144, 5, 211, 238, 71, 55, 45, 115, 89, 191, 114, 93, 177, 211,
			231, 144, 67, 38, 10, 13, 176, 124, 42, 235, 140, 208, 87, 203,
			31, 93, 38, 53, 162, 155, 154, 101, 94, 207, 149, 180, 194, 29,
			229, 43, 101, 141, 33, 229, 220, 63, 160, 179, 78, 147, 78, 18,
			234, 97, 143, 56, 137, 79, 19, 245, 184, 158, 191, 74, 238, 16,
			211, 212, 208, 97, 55, 244, 197, 66, 9, 158, 12, 232, 228, 67,
			98, 182, 150, 67, 86, 68, 129, 37, 184, 53, 20, 124, 77, 184,
			246, 134, 126, 253, 154, 128, 168, 9, 215, 222, 80, 240, 53, 225,
			218, 27, 167, 206, 103, 148, 97, 25, 55, 172, 75, 228, 107, 33,
			79, 179, 140, 155, 122, 25, 99, 109, 114, 92, 84, 81, 213, 79,
			40, 31, 196, 148, 243, 147, 97, 38, 60, 61, 145, 139, 30, 186,
			169, 223, 88, 84, 188, 181, 121, 100, 183, 156, 81, 200, 252, 234,
			205, 140, 50, 44, 227, 230, 74, 137, 84, 133, 92, 221, 50, 86,
			244, 139, 133, 34, 236, 188, 10, 71, 35, 26, 192, 175, 56, 139,
			101, 20, 189, 29, 213, 82, 18, 58, 104, 69, 191, 89, 86, 220,
			208, 65, 43, 170, 61, 213, 132, 131, 86, 242, 103, 50, 202, 176,
			140, 149, 243, 23, 72, 149, 232, 166, 110, 153, 159, 231, 110, 107,
			133, 79, 79, 108, 166, 62, 75, 176, 33, 205, 190, 140, 138, 36,
			162, 220, 129, 172, 62, 207, 95, 21, 251, 71, 71, 119, 84, 116,
			217, 240, 233, 194, 200, 21, 253, 115, 105, 100, 93, 24, 185, 162,
			246, 143, 46, 140, 92, 89, 56, 159, 81, 134, 101, 84, 172, 75,
			138, 139, 102, 25, 171, 250, 146, 226, 130, 38, 91, 213, 43, 139,
			106, 38, 6, 245, 234, 132, 11, 154, 108, 117, 225, 66, 70, 25,
			150, 177, 122, 105, 81, 113, 209, 45, 195, 214, 63, 82, 92, 208,
			28, 182, 190, 186, 164, 102, 162, 57, 108, 101, 14, 93, 8, 180,
			243, 23, 51, 202, 176, 12, 123, 113, 73, 236, 101, 65, 108, 252,
			237, 247, 178, 174, 27, 166, 101, 110, 232, 246, 71, 10, 148, 49,
			135, 56, 230, 51, 74, 179, 140, 13, 181, 151, 117, 221, 64, 140,
			87, 175, 145, 123, 68, 55, 13, 203, 188, 155, 187, 175, 21, 42,
			239, 223, 96, 127, 194, 145, 200, 243, 110, 254, 10, 249, 140, 152,
			166, 129, 142, 188, 167, 95, 44, 92, 129, 199, 111, 199, 89, 182,
			6, 97, 26, 194, 199, 247, 244, 187, 50, 132, 13, 225, 227, 123,
			202, 174, 134, 240, 241, 61, 21, 102, 134, 240, 241, 189, 243, 23,
			200, 38, 209, 77, 211, 50, 191, 202, 61, 214, 10, 246, 59, 97,
			54, 28, 71, 105, 56, 138, 148, 28, 96, 120, 178, 84, 29, 155,
			2, 106, 106, 150, 241, 85, 254, 19, 82, 198, 231, 5, 203, 248,
			250, 244, 217, 226, 39, 112, 228, 115, 249, 95, 0, 144, 73, 177,
			55, 78, 56, 75, 138, 208, 15, 105, 20, 16, 17, 16, 230, 66,
			14, 231, 74, 52, 230, 2, 98, 155, 161, 116, 73, 137, 137, 168,
			252, 55, 42, 138, 77, 161, 225, 55, 250, 215, 50, 211, 155, 66,
			195, 111, 84, 252, 153, 66, 195, 111, 84, 20, 155, 66, 195, 111,
			84, 20, 139, 141, 241, 173, 138, 98, 83, 68, 241, 183, 250, 55,
			139, 106, 38, 70, 241, 183, 19, 46, 24, 197, 223, 170, 40, 54,
			69, 20, 127, 123, 105, 145, 56, 130, 139, 110, 25, 53, 221, 42,
			220, 131, 6, 182, 167, 242, 210, 130, 66, 60, 30, 238, 211, 4,
			93, 34, 219, 61, 28, 79, 168, 72, 218, 165, 128, 246, 253, 113,
			148, 110, 194, 198, 90, 89, 249, 201, 20, 241, 95, 211, 191, 93,
			82, 50, 48, 254, 107, 202, 79, 114, 195, 213, 242, 153, 134, 24,
			242, 181, 11, 23, 201, 151, 66, 190, 97, 25, 117, 125, 185, 240,
			51, 112, 101, 251, 11, 44, 62, 96, 152, 160, 149, 224, 169, 184,
			190, 31, 113, 58, 149, 136, 1, 92, 215, 107, 150, 226, 138, 1,
			92, 87, 1, 108, 138, 0, 174, 159, 250, 40, 163, 80, 200, 149,
			130, 216, 113, 166, 110, 90, 198, 206, 223, 126, 199, 153, 34, 40,
			119, 244, 250, 178, 2, 101, 206, 33, 142, 12, 48, 6, 221, 142,
			218, 113, 38, 110, 51, 99, 71, 85, 79, 19, 109, 233, 254, 37,
			213, 211, 212, 231, 76, 203, 116, 245, 157, 235, 138, 221, 156, 224,
			144, 137, 154, 211, 44, 195, 85, 213, 211, 212, 231, 12, 203, 112,
			63, 186, 76, 182, 9, 34, 50, 119, 115, 191, 208, 10, 247, 62,
			188, 185, 79, 110, 31, 113, 211, 242, 246, 254, 65, 254, 187, 249,
			130, 216, 63, 115, 11, 150, 209, 252, 243, 246, 207, 28, 238, 159,
			166, 218, 49, 115, 98, 255, 204, 80, 186, 164, 154, 56, 17, 247,
			79, 91, 47, 21, 106, 19, 16, 242, 154, 55, 251, 15, 71, 106,
			52, 161, 28, 193, 198, 7, 242, 74, 97, 146, 91, 120, 150, 15,
			108, 34, 185, 227, 22, 52, 218, 250, 132, 154, 183, 140, 246, 233,
			171, 25, 165, 89, 70, 251, 147, 79, 51, 202, 176, 140, 246, 103,
			43, 217, 37, 194, 255, 4, 0, 0, 255, 255, 34, 155, 109, 97,
			247, 39, 0, 0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptor.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("milo.Buildbot")
	if err != nil {
		panic(err)
	}
	return ret
}
