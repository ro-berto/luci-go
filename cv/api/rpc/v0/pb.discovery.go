// Code generated by cproto. DO NOT EDIT.

package rpcpb

import "go.chromium.org/luci/grpc/discovery"

import "google.golang.org/protobuf/types/descriptorpb"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"cv.rpc.v0.Runs",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 236, 87, 221, 142, 27, 199,
			177, 230, 204, 80, 212, 110, 75, 218, 165, 122, 181, 130, 52, 134,
			143, 75, 43, 27, 250, 129, 118, 184, 187, 150, 113, 142, 44, 24,
			54, 197, 165, 124, 104, 105, 201, 61, 36, 215, 130, 13, 28, 72,
			205, 153, 38, 217, 209, 176, 123, 212, 211, 67, 106, 95, 32, 15,
			144, 155, 220, 230, 42, 175, 16, 36, 215, 1, 2, 228, 214, 65,
			174, 147, 23, 200, 35, 4, 65, 245, 204, 44, 185, 114, 126, 145,
			92, 154, 87, 172, 233, 234, 234, 174, 170, 175, 190, 234, 34, 63,
			189, 73, 26, 19, 21, 132, 83, 173, 102, 34, 155, 5, 74, 79,
			26, 113, 22, 138, 70, 56, 111, 176, 68, 52, 66, 53, 155, 41,
			217, 152, 239, 55, 116, 38, 131, 68, 43, 163, 232, 229, 112, 30,
			228, 223, 131, 249, 254, 206, 143, 93, 226, 245, 51, 73, 55, 136,
			43, 162, 27, 14, 56, 119, 215, 251, 174, 136, 232, 30, 169, 165,
			134, 153, 44, 189, 225, 130, 115, 119, 227, 224, 70, 176, 186, 45,
			232, 103, 50, 24, 216, 245, 126, 161, 71, 125, 178, 198, 231, 92,
			167, 66, 201, 27, 30, 56, 119, 189, 254, 153, 188, 243, 51, 135,
			212, 114, 117, 122, 157, 208, 193, 176, 57, 60, 25, 188, 60, 233,
			14, 142, 219, 173, 206, 211, 78, 251, 176, 94, 161, 151, 200, 197,
			227, 118, 247, 176, 211, 253, 178, 238, 160, 208, 63, 233, 118, 81,
			112, 169, 79, 174, 191, 104, 118, 134, 157, 238, 151, 47, 159, 246,
			250, 47, 7, 39, 79, 142, 58, 131, 65, 167, 215, 173, 87, 233,
			6, 33, 86, 30, 226, 114, 253, 2, 202, 237, 238, 97, 251, 240,
			229, 81, 115, 240, 172, 254, 5, 189, 66, 214, 7, 39, 173, 86,
			187, 125, 216, 62, 172, 55, 41, 33, 181, 167, 205, 206, 243, 246,
			97, 253, 9, 46, 181, 154, 221, 86, 251, 57, 138, 173, 39, 123,
			223, 6, 255, 84, 36, 31, 231, 255, 146, 209, 87, 63, 185, 78,
			106, 180, 186, 81, 57, 118, 200, 175, 171, 196, 185, 76, 189, 141,
			10, 61, 248, 101, 21, 90, 42, 57, 213, 98, 50, 53, 112, 176,
			119, 176, 15, 195, 41, 135, 231, 39, 173, 14, 52, 51, 51, 85,
			58, 13, 8, 129, 231, 34, 228, 50, 229, 17, 100, 50, 226, 26,
			204, 148, 67, 51, 97, 33, 106, 230, 43, 15, 224, 235, 60, 118,
			112, 16, 236, 193, 93, 84, 216, 41, 150, 118, 238, 61, 38, 112,
			170, 50, 152, 177, 83, 144, 202, 64, 150, 114, 48, 83, 145, 194,
			88, 196, 28, 248, 219, 144, 39, 6, 132, 132, 80, 205, 146, 88,
			48, 25, 114, 88, 8, 51, 181, 167, 20, 54, 2, 2, 223, 20,
			22, 212, 200, 48, 33, 129, 65, 168, 146, 83, 80, 227, 85, 53,
			96, 134, 16, 176, 191, 169, 49, 201, 167, 141, 198, 98, 177, 8,
			152, 189, 105, 30, 162, 92, 47, 109, 60, 239, 180, 218, 221, 65,
			123, 247, 32, 216, 35, 4, 78, 100, 204, 211, 20, 52, 127, 147,
			9, 205, 35, 24, 157, 2, 75, 146, 88, 132, 108, 20, 115, 136,
			217, 2, 148, 6, 54, 209, 156, 71, 96, 20, 222, 117, 161, 133,
			17, 114, 242, 0, 82, 53, 54, 11, 166, 57, 129, 72, 164, 70,
			139, 81, 102, 206, 133, 169, 188, 153, 72, 207, 41, 40, 9, 76,
			194, 78, 115, 0, 157, 193, 14, 60, 105, 14, 58, 131, 7, 4,
			94, 116, 134, 255, 219, 59, 25, 194, 139, 102, 191, 223, 236, 14,
			59, 237, 1, 244, 250, 208, 234, 117, 15, 59, 195, 78, 175, 59,
			128, 222, 83, 104, 118, 191, 129, 103, 157, 238, 225, 3, 224, 194,
			76, 185, 6, 254, 54, 209, 120, 123, 165, 65, 96, 0, 121, 20,
			16, 24, 112, 126, 238, 248, 177, 202, 175, 147, 38, 60, 20, 99,
			17, 66, 204, 228, 36, 99, 19, 14, 19, 53, 231, 90, 10, 57,
			129, 132, 235, 153, 72, 49, 137, 41, 48, 25, 17, 136, 197, 76,
			24, 102, 236, 135, 239, 121, 20, 16, 178, 70, 28, 151, 122, 245,
			202, 54, 254, 91, 163, 30, 173, 180, 201, 58, 113, 215, 46, 229,
			127, 79, 136, 91, 173, 208, 234, 118, 229, 216, 241, 31, 65, 63,
			147, 32, 100, 24, 103, 17, 79, 173, 157, 169, 152, 76, 119, 99,
			62, 231, 49, 8, 57, 86, 122, 102, 79, 2, 54, 82, 153, 1,
			6, 173, 175, 113, 75, 64, 118, 54, 160, 203, 223, 26, 48, 108,
			242, 41, 60, 12, 8, 33, 196, 171, 86, 28, 234, 109, 175, 93,
			34, 255, 67, 170, 213, 10, 30, 114, 211, 61, 244, 252, 251, 144,
			23, 45, 68, 60, 13, 181, 24, 21, 7, 229, 5, 143, 80, 89,
			90, 37, 151, 201, 5, 220, 233, 80, 239, 230, 197, 43, 228, 67,
			82, 67, 201, 173, 80, 239, 189, 234, 123, 254, 54, 156, 200, 34,
			82, 60, 42, 246, 7, 132, 108, 144, 139, 185, 150, 131, 106, 215,
			151, 178, 75, 189, 247, 110, 250, 228, 183, 78, 97, 198, 161, 222,
			173, 106, 221, 255, 149, 147, 123, 157, 66, 194, 101, 132, 33, 54,
			10, 173, 105, 131, 53, 213, 49, 184, 82, 36, 113, 196, 67, 134,
			117, 129, 250, 71, 76, 178, 9, 215, 48, 101, 169, 188, 99, 32,
			209, 42, 228, 41, 22, 31, 58, 51, 192, 237, 237, 57, 151, 6,
			78, 185, 1, 165, 137, 253, 220, 207, 100, 111, 33, 243, 77, 192,
			223, 78, 89, 150, 34, 204, 88, 28, 219, 229, 55, 153, 50, 12,
			147, 10, 11, 102, 129, 107, 1, 33, 249, 162, 88, 49, 138, 192,
			136, 3, 155, 51, 17, 35, 230, 87, 156, 117, 172, 51, 151, 150,
			178, 75, 189, 91, 27, 155, 4, 10, 95, 93, 234, 221, 174, 214,
			253, 171, 165, 171, 58, 147, 136, 166, 21, 11, 24, 142, 219, 43,
			22, 236, 150, 141, 77, 242, 103, 183, 48, 225, 81, 47, 168, 126,
			224, 255, 201, 45, 109, 172, 94, 50, 205, 70, 5, 42, 49, 104,
			133, 130, 144, 57, 131, 20, 169, 21, 99, 80, 146, 151, 100, 48,
			86, 113, 172, 22, 184, 63, 13, 185, 100, 90, 40, 220, 98, 116,
			198, 63, 69, 114, 216, 15, 96, 168, 185, 173, 200, 48, 86, 24,
			87, 102, 236, 62, 35, 102, 121, 2, 152, 49, 124, 150, 152, 212,
			166, 11, 143, 55, 1, 110, 60, 8, 144, 27, 181, 221, 201, 164,
			178, 121, 203, 81, 157, 195, 140, 205, 10, 226, 60, 214, 234, 71,
			60, 68, 163, 204, 230, 56, 204, 180, 230, 210, 196, 167, 5, 55,
			229, 54, 141, 141, 18, 0, 124, 108, 237, 174, 56, 106, 195, 200,
			12, 207, 11, 16, 47, 24, 134, 74, 151, 0, 50, 165, 174, 1,
			149, 216, 146, 17, 178, 48, 92, 30, 220, 82, 114, 44, 48, 5,
			48, 92, 141, 82, 10, 33, 178, 107, 252, 110, 138, 61, 7, 19,
			224, 47, 101, 151, 122, 193, 251, 255, 69, 126, 83, 226, 185, 74,
			189, 135, 213, 45, 255, 23, 103, 120, 94, 245, 128, 64, 211, 126,
			14, 25, 162, 117, 196, 203, 67, 120, 4, 139, 41, 50, 252, 138,
			114, 161, 138, 60, 110, 52, 147, 169, 176, 247, 31, 107, 53, 35,
			231, 18, 106, 84, 89, 24, 175, 254, 122, 47, 125, 117, 86, 213,
			26, 24, 72, 37, 119, 207, 78, 37, 96, 144, 200, 36, 139, 191,
			95, 185, 85, 7, 61, 217, 88, 202, 46, 245, 30, 94, 165, 228,
			231, 37, 20, 47, 80, 239, 179, 234, 53, 255, 13, 44, 219, 50,
			250, 131, 110, 101, 22, 42, 41, 48, 24, 9, 51, 99, 233, 107,
			188, 101, 56, 229, 225, 107, 196, 31, 179, 142, 217, 234, 147, 145,
			165, 96, 27, 250, 163, 147, 193, 16, 186, 189, 225, 170, 129, 119,
			57, 201, 18, 210, 65, 27, 218, 50, 194, 15, 232, 204, 59, 30,
			60, 206, 237, 76, 217, 156, 195, 156, 197, 25, 7, 219, 175, 204,
			148, 73, 120, 181, 188, 233, 171, 128, 28, 252, 127, 227, 63, 247,
			35, 48, 44, 239, 145, 19, 235, 50, 144, 23, 28, 12, 212, 50,
			144, 23, 92, 234, 125, 118, 117, 139, 220, 41, 226, 88, 163, 222,
			231, 213, 45, 255, 134, 141, 10, 151, 17, 66, 38, 68, 26, 27,
			103, 113, 124, 186, 146, 145, 154, 131, 154, 87, 150, 178, 75, 189,
			207, 235, 148, 220, 43, 12, 93, 164, 94, 179, 90, 247, 253, 165,
			161, 76, 254, 13, 83, 23, 29, 212, 93, 145, 93, 234, 53, 175,
			108, 146, 157, 194, 212, 26, 245, 90, 213, 45, 127, 171, 4, 241,
			25, 98, 86, 108, 172, 57, 168, 180, 188, 206, 154, 75, 189, 86,
			157, 146, 63, 58, 216, 105, 176, 63, 60, 115, 235, 254, 119, 142,
			173, 216, 76, 190, 17, 25, 135, 206, 97, 201, 60, 121, 107, 177,
			107, 121, 67, 179, 9, 150, 168, 33, 82, 216, 73, 242, 242, 76,
			27, 31, 226, 83, 109, 183, 16, 241, 173, 155, 54, 62, 20, 209,
			206, 3, 88, 32, 191, 16, 216, 133, 85, 5, 203, 94, 83, 14,
			18, 9, 166, 124, 240, 32, 207, 36, 103, 60, 147, 243, 214, 136,
			199, 74, 78, 82, 203, 230, 187, 32, 162, 156, 168, 64, 37, 236,
			77, 198, 225, 53, 63, 133, 76, 10, 252, 91, 48, 214, 170, 145,
			178, 37, 186, 149, 11, 232, 227, 90, 41, 57, 212, 123, 182, 126,
			169, 148, 60, 234, 61, 219, 216, 36, 183, 108, 48, 28, 234, 29,
			185, 215, 252, 107, 43, 96, 94, 6, 33, 223, 224, 212, 80, 167,
			52, 134, 173, 228, 104, 125, 179, 148, 60, 234, 29, 209, 45, 242,
			127, 214, 152, 75, 189, 158, 187, 237, 31, 66, 249, 8, 47, 189,
			230, 210, 8, 115, 10, 197, 87, 12, 145, 8, 167, 184, 56, 83,
			82, 25, 37, 69, 200, 226, 248, 20, 159, 22, 154, 179, 52, 103,
			164, 252, 0, 44, 230, 158, 123, 177, 148, 28, 234, 245, 214, 234,
			165, 228, 81, 175, 183, 117, 109, 84, 179, 67, 198, 199, 228, 247,
			151, 200, 238, 223, 123, 79, 235, 36, 108, 204, 247, 108, 170, 138,
			185, 100, 61, 156, 7, 58, 9, 131, 249, 158, 255, 175, 206, 52,
			59, 31, 144, 43, 95, 114, 211, 207, 100, 159, 191, 201, 120, 106,
			222, 29, 103, 14, 62, 35, 213, 126, 38, 83, 250, 9, 169, 229,
			138, 212, 14, 52, 249, 121, 193, 185, 189, 254, 213, 239, 141, 58,
			79, 238, 127, 123, 247, 31, 59, 243, 88, 39, 97, 50, 250, 234,
			119, 235, 56, 22, 108, 86, 238, 159, 141, 5, 155, 63, 140, 5,
			63, 140, 5, 255, 246, 88, 112, 181, 66, 139, 177, 96, 171, 242,
			121, 57, 22, 20, 127, 189, 10, 245, 182, 43, 143, 200, 31, 92,
			226, 214, 42, 180, 10, 149, 143, 28, 255, 59, 251, 252, 75, 33,
			229, 122, 46, 66, 68, 64, 162, 82, 158, 22, 79, 118, 123, 54,
			30, 36, 236, 147, 112, 87, 243, 84, 101, 58, 228, 22, 135, 183,
			240, 135, 225, 193, 17, 60, 151, 48, 223, 247, 225, 196, 38, 31,
			97, 166, 65, 45, 36, 104, 145, 190, 14, 242, 165, 23, 8, 167,
			24, 59, 172, 74, 32, 205, 146, 68, 105, 251, 230, 180, 16, 156,
			239, 65, 243, 184, 99, 241, 134, 211, 136, 84, 70, 132, 188, 216,
			216, 85, 48, 98, 225, 235, 5, 211, 81, 106, 193, 201, 140, 24,
			137, 24, 73, 106, 146, 49, 205, 164, 225, 54, 214, 168, 123, 28,
			115, 134, 85, 16, 42, 105, 24, 62, 204, 190, 134, 25, 19, 18,
			161, 202, 117, 138, 87, 179, 68, 207, 229, 228, 11, 24, 241, 177,
			210, 248, 74, 56, 187, 133, 141, 54, 254, 22, 220, 98, 60, 209,
			106, 46, 34, 14, 44, 138, 236, 195, 137, 197, 203, 19, 237, 155,
			233, 84, 101, 13, 235, 107, 17, 195, 124, 90, 170, 33, 137, 195,
			218, 101, 114, 159, 84, 107, 182, 135, 221, 118, 255, 219, 127, 31,
			114, 26, 1, 205, 77, 166, 101, 106, 59, 72, 196, 13, 19, 113,
			154, 83, 104, 45, 167, 255, 219, 181, 203, 165, 132, 15, 247, 43,
			126, 41, 121, 212, 187, 253, 209, 39, 132, 228, 99, 222, 29, 100,
			144, 114, 54, 187, 179, 182, 77, 250, 101, 199, 188, 231, 214, 253,
			118, 209, 3, 243, 100, 66, 7, 113, 26, 76, 130, 149, 158, 88,
			178, 85, 222, 14, 31, 61, 122, 180, 127, 240, 241, 195, 253, 131,
			189, 221, 253, 221, 17, 139, 66, 54, 230, 59, 171, 45, 234, 222,
			185, 22, 117, 239, 92, 139, 186, 183, 177, 89, 18, 251, 95, 2,
			0, 0, 255, 255, 202, 211, 227, 146, 115, 18, 0, 0},
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
	ret, err := discovery.GetDescriptorSet("cv.rpc.v0.Runs")
	if err != nil {
		panic(err)
	}
	return ret
}
