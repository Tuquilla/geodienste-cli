package stac

var collections_testdata = `{
  "collections": [
    {
      "stac_version": "1.0.0",
      "type": "Collection",
      "id": "klaeranlagen_mit_finanzkennzahlen",
      "title": "Abwasserreinigungsanlagen (ARA) mit Finanzkennzahlen",
      "description": "Die erhobenen Daten liefern eine Übersicht über den Stand und die Entwicklung der Abwasserentsorgung in der Schweiz. Sie werden von den Behörden und Fachstellen aller Staatsebenen, von spezialisierten Ingenieur- und Planungsunternehmen, den Fachverbänden, den Anlagebetreibern sowie der Forschung für folgende Ziele verwendet:\u003C/p\u003E\n\u003Cul\u003E\n  \u003Cli\u003EErfolgs- und Leistungskontrolle der Abwasserreinigung\u003C/li\u003E\n  \u003Cli\u003EFrüherkennung von zukünftigen Herausforderungen im Bereich der Abwasserreinigung\u003C/li\u003E\n  \u003Cli\u003EErarbeitung von Strategien Massnahmen, um vorhandenen Defiziten und zukünftigen Herausforderungen zu begegnen\u003C/li\u003E\n  \u003Cli\u003EEinhaltung internationaler Verpflichtungen und Abkommen zur Datenlieferung\u003C/li\u003E\n\u003C/ul\u003E\n\u003Cp\u003EDie Erhebung wird periodisch durchgeführt. Die ARA-Daten beschreiben die technischen und betrieblichen Aspekte sowie die Leistungen der Abwasserreinigungsanlagen. Die Daten basieren auf dem minimalen Geodatenmodell (MGDM) «Kläranlagendatenbank (ARA-DB)» und «Regionale Entwässerungsplanung» (\u003Ca href='https://geobasisdaten.ch/detail/817381/' target='_blank'\u003EID 134.5\u003C/a\u003E, \u003Ca href='https://geobasisdaten.ch/detail/819149/' target='_blank'\u003EID 128.1\u003C/a\u003E, Version 1.2).\u003C/p\u003E\n\u003Cbr\u003E\n\u003Cp\u003EDieses Angebot beinhaltet Finanzkennzahlen zu den Kläranlagen (ARA). Der Zugang ist deshalb eingeschränkt. Ein Angebot \u003Ca href=\"klaeranlagen_ohne_finanzkennzahlen\"\u003E ohne Finanzkennzahlen\u003C/a\u003E steht ebenfalls zur Verfügung.\u003C/p\u003E\n",
      "keywords": [
        "Kläranlagendatenbank",
        "Abwasserreinigungsanlagen (ARA)",
        "Abwasserreinigungsanlagen",
        "Abwasserbehandlung",
        "Kläranlagen",
        "Abwasserentsorgung",
        "Wasseraufbereitung",
        "Abwassertechnik",
        "Abwasserlösungen",
        "Umwelttechnik",
        "Abwasserfilter",
        "Abwasserreinigung",
        "ID 134.5"
      ],
      "license": "proprietary",
      "extent": {
        "spatial": {
          "bbox": [
            [5.9559019, 45.8179584, 10.4921713, 47.80845431]
          ]
        },
        "temporal": {
          "interval": [
            [
              "2025-09-24T06:39:03Z",
              null]
          ]
        }
      },
      "providers": [
        {
          "name": "Kanton AR",
          "roles": [
            "licensor",
            "producer"
          ]
        },
        {
          "name": "Kanton BL",
          "roles": [
            "licensor",
            "producer"
          ]
        },
        {
          "name": "Kanton SG",
          "roles": [
            "licensor",
            "producer"
          ]
        },
        {
          "name": "Kanton SH",
          "roles": [
            "licensor",
            "producer"
          ]
        },
        {
          "name": "Kanton SZ",
          "roles": [
            "licensor",
            "producer"
          ]
        },
        {
          "name": "geodienste.ch KGK-CGC",
          "roles": [
            "processor",
            "host"
          ],
          "url": "https://geodienste.ch/services/klaeranlagen_mit_finanzkennzahlen"
        }
      ],
      "links": [
        {
          "href": "https://geodienste.ch/stac/collections/klaeranlagen_mit_finanzkennzahlen",
          "rel": "self",
          "type": "application/json",
          "title": "This collection"
        },
        {
          "href": "https://geodienste.ch/stac",
          "rel": "root",
          "type": "application/json"
        },
        {
          "href": "https://geodienste.ch/stac",
          "rel": "parent",
          "type": "application/json"
        },
        {
          "href": "https://geodienste.ch/stac/collections/klaeranlagen_mit_finanzkennzahlen/items",
          "rel": "items",
          "type": "application/geo+json",
          "title": "Information about the collection items"
        },
        {
          "href": "https://geodienste.ch/services/klaeranlagen_mit_finanzkennzahlen",
          "rel": "license",
          "type": "text/html",
          "title": "Terms of use for providers"
        },
        {
          "href": "https://www.geocat.ch/geonetwork/srv/ger/catalog.search#/metadata/6f192077-f4e5-4ab7-80dc-4e5affb5c1dd",
          "rel": "describedby",
          "type": "text/html",
          "title": "Metadata"
        },
        {
          "href": "https://geodienste.ch/services/klaeranlagen_mit_finanzkennzahlen",
          "rel": "related",
          "type": "text/html",
          "title": "Service on geodienste.ch"
        }
      ]
    }
	]
}
`
