// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package lockss

// Purpose: Hold content that I will probably prune at some point.

/*
func NewConfig(xmlDoc *xmlquery.Node) (LOCKSSConfig, error) {

	proxyAccessIPInclude, err := ProxyAccessIPInclude(xmlDoc)
	if err != nil {
		return LOCKSSConfig{}, err
	}

	proxyAccessLogForbidden, err := ProxyAccessIPLogForbidden(xmlDoc)
	if err != nil {
		return LOCKSSConfig{}, err
	}

	proxyPort, err := ProxyPort(xmlDoc)
	if err != nil {
		return LOCKSSConfig{}, err
	}

	proxyAuditPort, err := ProxyAuditPort(xmlDoc)
	if err != nil {
		return LOCKSSConfig{}, err
	}

	noManifestIndexResponses, err := ProxyNoManifestIndexResponses(xmlDoc)
	if err != nil {
		return LOCKSSConfig{}, err
	}

	cfg := LOCKSSConfig{
		Proxy: Proxy{
			NoManifestIndexResponses: noManifestIndexResponses,
			Port:                     proxyPort,
			AuditPort:                proxyAuditPort,
			Access: ProxyAccess{
				IPInclude:      proxyAccessIPInclude,
				IPLogForbidden: proxyAccessLogForbidden,
			},
		},
	}

	return cfg, nil

}

*/

// Content pulled from n2n "main":

//fmt.Printf("Parsed document: %+v\n", doc)

// appExitOnce := xmlquery.Find(doc, "//property//property[@name='app.exitOnce']")
// fmt.Printf("Find results = Type: %T, Value: %+v\n", appExitOnce, appExitOnce)

// fmt.Println(appExitOnce[0].Attr)

// proxyLogForbidden := xmlquery.Find(doc, "//property//property[@name='proxy']//property[@name='ip.logForbidden']")
// //fmt.Printf("Find results = Type: %T, Value: %+v\n", proxyLogForbidden, proxyLogForbidden)

// for _, item := range proxyLogForbidden {
// 	fmt.Printf("%+v\n", item)
// }

// ipInclude := xmlquery.Find(doc, "//property//property[@name='proxy']//property[@name='ip.include']")
// ipInclude, err := xmlquery.QueryAll(doc, "//property//property[@name='proxy']//property[@name='ip.include']//list//value")
// if err != nil {
// 	fmt.Printf("Error occurred running XPath query: %v\n", err)
// 	os.Exit(1)
// }
// xpathExpression, err := xpath.Compile("//property[@name='proxy']//property[@name='access']//property[@name='ip.include']//list/value")
// if err != nil {
// 	fmt.Printf("Error occurred compiling XPath query: %v\n", err)
// 	os.Exit(1)
// }

//fmt.Printf("Find results = Type: %T, Value: %+v\n", ipInclude, ipInclude)

// for _, item := range ipInclude {
// 	fmt.Printf("%+v\n", item)
// }

//fmt.Printf("%+v\n", ipInclude.SelectElements("property"))

/*

func ProxyAccessIPInclude(xmlDoc *xmlquery.Node) ([]string, error) {

	myFuncName := caller.GetFuncName()

	includedIPs := make([]string, 0, 5)

	xpathExpression := "//property[@name='proxy']//property[@name='access']//property[@name='ip.include']//list/value"

	xmlqueryNodes, err := xmlquery.QueryAll(xmlDoc, xpathExpression)
	if err != nil {
		return nil, fmt.Errorf("error occurred running XPath query: %w", err)
	}

	// If the query failed to find a match then this will be nil. We should
	// halt further attempts to retrieve more specific values.
	if xmlqueryNodes == nil {
		return nil, fmt.Errorf(
			"%s: unable to query value using xpathExpression: %q",
			myFuncName,
			xpathExpression,
		)
	}

	for _, prop := range xmlqueryNodes {
		// fmt.Printf("Raw XML: %q\n", prop.OutputXML(true))
		// fmt.Printf("Inner Text: %q\n", prop.InnerText())
		includedIPs = append(includedIPs, prop.InnerText())
	}

	return includedIPs, nil

}

func ProxyAccessIPLogForbidden(xmlDoc *xmlquery.Node) (bool, error) {

	myFuncName := caller.GetFuncName()

	xpathExpression := "//property[@name='proxy']//property[@name='access']//property[@name='ip.logForbidden']"

	xmlqueryNode, err := xmlquery.Query(xmlDoc, xpathExpression)
	if err != nil {
		return false, fmt.Errorf("error occurred running XPath query: %w", err)
	}

	// If the query failed to find a match then this will be nil. We should
	// halt further attempts to retrieve more specific values.
	if xmlqueryNode == nil {
		return false, fmt.Errorf(
			"%s: unable to query value using xpathExpression: %q",
			myFuncName,
			xpathExpression,
		)
	}

	// fmt.Printf("Raw: %+v\n", xmlqueryNode)
	// fmt.Printf("XML: %+v\n", xmlqueryNode.OutputXML(true))
	// fmt.Printf("Inner Text: %+v\n", xmlqueryNode.InnerText())
	//fmt.Printf("SelectAttr: %+v\n", xmlqueryNode.SelectAttr("value"))

	logForbiddenStr := xmlqueryNode.SelectAttr("value")
	logForbiddenBool, err := strconv.ParseBool(logForbiddenStr)
	if err != nil {
		return false, err
	}

	return logForbiddenBool, nil

}

func ProxyPort(xmlDoc *xmlquery.Node) (int, error) {

	xpathExpression := "//property[@name='proxy']//property[@name='port']"

	xmlqueryNode, err := xmlquery.Query(xmlDoc, xpathExpression)
	if err != nil {
		return 0, fmt.Errorf("error occurred running XPath query: %w", err)
	}

	// fmt.Printf("Raw: %+v\n", xmlqueryNode)
	// fmt.Printf("XML: %+v\n", xmlqueryNode.OutputXML(true))
	// fmt.Printf("Inner Text: %+v\n", xmlqueryNode.InnerText())
	//fmt.Printf("SelectAttr: %+v\n", xmlqueryNode.SelectAttr("value"))

	proxyPortStr := xmlqueryNode.SelectAttr("value")
	// FIXME: TCP ports are between 0-65535. Will 16 bits do what I need here?
	proxyPortInt, err := strconv.ParseInt(proxyPortStr, 10, 16)
	if err != nil {
		return 0, err
	}

	// FIXME: Convert to plain int?
	return int(proxyPortInt), nil

}

func ProxyAuditPort(xmlDoc *xmlquery.Node) (int, error) {

	xpathExpression := "//property[@name='proxy']//property[@name='audit.port']"

	xmlqueryNode, err := xmlquery.Query(xmlDoc, xpathExpression)
	if err != nil {
		return 0, fmt.Errorf("error occurred running XPath query: %w", err)
	}

	// fmt.Printf("Raw: %+v\n", xmlqueryNode)
	// fmt.Printf("XML: %+v\n", xmlqueryNode.OutputXML(true))
	// fmt.Printf("Inner Text: %+v\n", xmlqueryNode.InnerText())
	//fmt.Printf("SelectAttr: %+v\n", xmlqueryNode.SelectAttr("value"))

	portStr := xmlqueryNode.SelectAttr("value")
	// FIXME: TCP ports are between 0-65535. Will 16 bits do what I need here?
	portInt, err := strconv.ParseInt(portStr, 10, 16)
	if err != nil {
		return 0, err
	}

	// FIXME: Convert to plain int?
	return int(portInt), nil

}

func ProxyNoManifestIndexResponses(xmlDoc *xmlquery.Node) (string, error) {

	xpathExpression := "//property[@name='proxy']//property[@name='noManifestIndexResponses']"

	xmlqueryNode, err := xmlquery.Query(xmlDoc, xpathExpression)
	if err != nil {
		return "", fmt.Errorf("error occurred running XPath query: %w", err)
	}

	// fmt.Printf("Raw: %+v\n", xmlqueryNode)
	// fmt.Printf("XML: %+v\n", xmlqueryNode.OutputXML(true))
	// fmt.Printf("Inner Text: %+v\n", xmlqueryNode.InnerText())
	//fmt.Printf("SelectAttr: %+v\n", xmlqueryNode.SelectAttr("value"))

	noManifestIndexResponses := xmlqueryNode.SelectAttr("value")

	return noManifestIndexResponses, nil

}

*/
