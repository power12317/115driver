package driver

const (
	UADefalut    = "Mozilla/5.0"
	UA115Browser = "Mozilla/5.0 115Browser/23.9.3.2"
	UA115Disk    = "Mozilla/5.0 115disk/30.1.0"
	UA115Desktop = "Mozilla/5.0; Windows NT/10.0.20348; 115Desktop/2.0.6.6"
	UAIosApp     = "Mozilla/5.0; Darwin/10.0; UDown/30.1.0"
)

const (
	CookieDomain115 = ".115.com"

	CookieUrl = "https://115.com"

	CookieNameUid  = "UID"
	CookieNameCid  = "CID"
	CookieNameSeid = "SEID"
)

const (
	OSSRegionID = "oss-cn-shenzhen"
	OSSEndpoint = "cn-shenzhen.oss.aliyuncs.com" // 双栈域名

	OSSUserAgent               = "aliyun-sdk-android/2.9.1"
	OssSecurityTokenHeaderName = "X-OSS-Security-Token"
)

const (
	KB = 1 << (10 * (iota + 1))
	MB
	GB
)
