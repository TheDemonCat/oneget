package downloader

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHtmlParser_ParseTotalReleases(t *testing.T) {
	body := strings.NewReader(releasesTotal)
	parser, err := NewHtmlParser()

	if err != nil {
		t.Fatal(err)
	}

	releases, err := parser.ParseTotalReleases(body)
	assert.NoError(t, err)
	assert.Lenf(t, releases, 69, "releases count must be equal")

}

func TestHtmlParser_ParseProjectReleases(t *testing.T) {
	body := strings.NewReader(projectReleases)
	parser, err := NewHtmlParser()

	if err != nil {
		t.Fatal(err)
	}

	releases, err := parser.ParseProjectReleases(body)
	if err != nil {
		return
	}

	assert.NoError(t, err)
	assert.Lenf(t, releases, 165, "releases count must be equal")

}

func TestHtmlParser_ParseProjectRelease(t *testing.T) {
	body := strings.NewReader(projectReleaseHtml)
	parser, err := NewHtmlParser()

	if err != nil {
		t.Fatal(err)
	}

	releaseFiles, err := parser.ParseProjectRelease(body)
	if err != nil {
		return
	}

	assert.NoError(t, err)
	assert.Lenf(t, releaseFiles, 24, "release files count must be equal")

}

func TestHtmlParser_ParseReleaseFiles(t *testing.T) {
	body := strings.NewReader(releaseFiles)
	parser, err := NewHtmlParser()

	if err != nil {
		t.Fatal(err)
	}

	fileUrls, err := parser.ParseReleaseFiles(body)
	if err != nil {
		return
	}

	assert.NoError(t, err)
	assert.Lenf(t, fileUrls, 3, "file urls count must be equal")

}

var releaseFiles = `
<table class="common-container">
   <tbody>
      <tr class="header-container">
         <td>
            <div class="navbar">
               <div class="header">
                  <!-- Settings toolbar contains location and language components and so on -->
                  <div class="settings-toolbar">
                     <li id="countryDropdown" class="dropdown">
                        <a class="dropdown-toggle" data-toggle="dropdown" href="javascript:void(0);">
                           <div class="countryFlagImg flag-RU"></div>
                           <span class="textValue">Россия</span>
                           <b class="caret"></b>
                        </a>
                        <ul class="dropdown-menu localizationDropDownMenu">
                           <li data-value="RU"><a href="javascript:void(0);" data-value="RU"><img class="countryFlagImg flag-RU"><span>Россия</span></a></li>
                           <li data-value="UA"><a href="javascript:void(0);" data-value="UA"><img class="countryFlagImg flag-UA"><span>Украина</span></a></li>
                           <li data-value="LV"><a href="javascript:void(0);" data-value="LV"><img class="countryFlagImg flag-LV"><span>Латвия</span></a></li>
                           <li data-value="BY"><a href="javascript:void(0);" data-value="BY"><img class="countryFlagImg flag-BY"><span>Беларусь</span></a></li>
                           <li data-value="GE"><a href="javascript:void(0);" data-value="GE"><img class="countryFlagImg flag-GE"><span>Грузия</span></a></li>
                           <li data-value="KG"><a href="javascript:void(0);" data-value="KG"><img class="countryFlagImg flag-KG"><span>Кыргызстан</span></a></li>
                           <li data-value="KZ"><a href="javascript:void(0);" data-value="KZ"><img class="countryFlagImg flag-KZ"><span>Казахстан</span></a></li>
                           <li data-value="EE"><a href="javascript:void(0);" data-value="EE"><img class="countryFlagImg flag-EE"><span>Эстония</span></a></li>
                           <li data-value="AM"><a href="javascript:void(0);" data-value="AM"><img class="countryFlagImg flag-AM"><span>Армения</span></a></li>
                           <li data-value="LT"><a href="javascript:void(0);" data-value="LT"><img class="countryFlagImg flag-LT"><span>Литва</span></a></li>
                           <li data-value="MD"><a href="javascript:void(0);" data-value="MD"><img class="countryFlagImg flag-MD"><span>Молдова</span></a></li>
                           <li data-value="UZ"><a href="javascript:void(0);" data-value="UZ"><img class="countryFlagImg flag-UZ"><span>Узбекистан</span></a></li>
                           <li data-value="AZ"><a href="javascript:void(0);" data-value="AZ"><img class="countryFlagImg flag-AZ"><span>Азербайджан</span></a></li>
                        </ul>
                     </li>
                     <li id="localeDropdown" class="dropdown">
                        <a class="dropdown-toggle" data-toggle="dropdown" href="javascript:void(0);">
                        <span class="textValue">Русский</span>
                        <b class="caret"></b>
                        </a>
                        <ul class="dropdown-menu localizationDropDownMenu">
                           <li data-value="ru"><a href="javascript:void(0);" data-value="ru"><span>Русский</span></a></li>
                        </ul>
                     </li>
                     <script type="application/javascript">
                        localization.initLocalization();
                     </script>
                     <span class="portal-name">
                     <a href="https://portal.1c.ru">
                     Портал 1С:ИТС
                     </a>
                     </span>
                     <span class="exit-btn">
                     <span class="exit-btn-icon"><img src="/resources/resources/img/exit-btn-icon.png"></span>
                     <a href=" /exit">
                     Выход
                     </a>
                     </span>
                     <span class="login">
                        <span class="login-icon"><img src="/resources/resources/img/login-icon.png"></span>
                        <a href="https://login.1c.ru/user/profile">
                           <nobr>prof-concept</nobr>
                        </a>
                     </span>
                  </div>
                  <!-- Service toolbar contains service name and menu toolbar-->
                  <div class="service-toolbar">
                     <div class="service-icon">
                        <span><img src="/resources/resources/img/releases-service-icon.png"></span>
                     </div>
                     <div class="service-info">
                        <div class="service-name">
                           <span>
                           1C:Обновление программ
                           </span>
                        </div>
                        <div class="menu-toolbar">
                           <span class="menu-main-page">
                           <img src="/resources/resources/img/main-icon.png">
                           <a href="/">
                           Главная
                           </a>
                           </span>
                           <span class="menu-news-page">
                           <img src="/resources/resources/img/news-icon.png">
                           <a href="/news">
                           Новости
                           </a>
                           </span>
                           <span class="menu-classifiers-page">
                           <span style="vertical-align:middle"><i class="fa fa-th-list" aria-hidden="false"></i></span>
                           <a href="/classifiers/total">
                           Классификаторы
                           </a>
                           </span>
                           <span class="menu-profile-page">
                           <img src="/resources/resources/img/profile-icon.png">
                           <a href="https://portal.1c.ru/software">
                           Личный кабинет
                           </a>
                           </span>
                           <span class="menu-about-page">
                           <img src="/resources/resources/img/about-icon.png">
                           <a href="https://portal.1c.ru/applications/4">
                           1C:Обновление программ
                           </a>
                           </span>
                        </div>
                     </div>
                  </div>
               </div>
            </div>
         </td>
      </tr>
      <tr class="body-container">
         <td style="vertical-align: top">
            <div class="body-content">
               <div class="downloadDist">
                  <a href="https://dl04.1c.ru/public/file/get/c1aa644a-592e-4b72-b07c-2ef7999ef497">Скачать дистрибутив</a><br><br>
                  <a href="https://dl05.1c.ru/public/file/get/c1aa644a-592e-4b72-b07c-2ef7999ef497">Скачать дистрибутив с зеркала 1</a><br>
                  <a href="https://dl03.1c.ru/public/file/get/c1aa644a-592e-4b72-b07c-2ef7999ef497">Скачать дистрибутив с зеркала 2</a><br>
               </div>
            </div>
         </td>
      </tr>
      <tr class="footer-container">
         <td>
            <footer class="footer">
               <div class="innerFooter">
                  <div class="footerInfo">
                     <span>
                     © 1С-Софт, 2021. Все права защищены
                     </span>
                  </div>
               </div>
            </footer>
         </td>
      </tr>
   </tbody>
</table>
`
var projectReleaseHtml = `
<div class="body-content">
   <div class="files-container">
      <div class="section-title">
         <a href="/project/Platform83">Технологическая платформа 8.3</a>,
         версия&nbsp;8.3.18.1363
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\1cv8upd_8_3_18_1363.htm">Технологическая платформа 8.3. Версия 8.3.18.1363. Список изменений и порядок обновления</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\setuptc_8_3_18_1363.rar">Тонкий клиент 1С:Предприятия для Windows</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\setuptc64_8_3_18_1363.rar">Тонкий клиент 1С:Предприятие (64-bit) для Windows</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\thin.client_8_3_18_1363.deb32.tar.gz">Тонкий клиент 1С:Предприятия для DEB-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\thin.client_8_3_18_1363.rpm32.tar.gz">Тонкий клиент 1С:Предприятия для RPM-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\thin.client_8_3_18_1363.deb64.tar.gz">Тонкий клиент 1С:Предприятия (64-bit) для DEB-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\thin.client_8_3_18_1363.rpm64.tar.gz">Тонкий клиент 1С:Предприятия (64-bit) для RPM-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\thin.osx_8_3_18_1363.dmg">Тонкий клиент 1С:Предприятия для macOS</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\windows_8_3_18_1363.rar">Технологическая платформа 1С:Предприятия для Windows</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\windows64full_8_3_18_1363.rar">Технологическая платформа 1С:Предприятия (64-bit) для Windows</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\client_8_3_18_1363.deb32.tar.gz">Клиент 1С:Предприятия для DEB-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\client_8_3_18_1363.rpm32.tar.gz">Клиент 1С:Предприятия для RPM-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\client_8_3_18_1363.deb64.tar.gz">Клиент 1С:Предприятия (64-bit) для DEB-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\client_8_3_18_1363.rpm64.tar.gz">Клиент 1С:Предприятия (64-bit) для RPM-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\deb_8_3_18_1363.tar.gz">Cервер 1С:Предприятия для DEB-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\clientosx_8_3_18_1363.dmg">Клиент 1С:Предприятия для macOS</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\rpm_8_3_18_1363.tar.gz">Cервер 1С:Предприятия для RPM-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\windows64_8_3_18_1363.rar">Cервер 1С:Предприятия (64-bit) для Windows</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\deb64_8_3_18_1363.tar.gz">Cервер 1С:Предприятия (64-bit) для DEB-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\rpm64_8_3_18_1363.tar.gz">Cервер 1С:Предприятия (64-bit) для RPM-based Linux-систем</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\addin_8_3_18_1363.zip">Технология внешних компонент</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\demo.zip">Демонстрационная информационная база</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\Collations.rar">Файл настройки сортировки для Oracle Database</a>
      </div>
      <div class="formLine">
         <a href="/version_file?nick=Platform83&amp;ver=8.3.18.1363&amp;path=Platform\8_3_18_1363\Err_Other.htm">Решение текущих проблем работы с различными СУБД и ОС</a>
      </div>
      <hr class="fileDelimiter">
      <div class="formLine">
         <p>
            <a href="https://bugboard.v8.1c.ru/version/plt8gen/000025703" target="_blank">
            <i class="fa fa-bug"></i>
            Проблемные ситуации и ошибки в версии 8.3.18.1363
            </a>
         </p>
      </div>
   </div>
</div>
`
var projectReleases = `
<div class="projects-container">
   <div>
      <span class="project-name">Технологическая платформа 8.3</span>
   </div>
   <div class="formLine">
      <p></p>
      <p><a href="https://bugboard.v8.1c.ru/catalog.html">Сервис публикации ошибок</a></p>
      <p>Доступ к публикуемым ошибкам технологической платформы "1С:Предприятие 8": каталог, поиск, сравнение версий, email-подписка. </p>
   </div>
   <div class="formLine">
      Обращаем ваше внимание, что с мая 2020 года фирма "1С" перешла на бесконтактную форму сопровождения, в связи с чем был прекращен выпуск дисков ИТС.<br>Подробнее см. информационное письмо № 27079 от 17.04.2020 <a href="https://1c.ru/news/info.jsp?id=27079" target="_blank">https://1c.ru/news/info.jsp?id=27079</a>.
   </div>
   <div class="formLine projectPageTitle">
      <h5>
         Обновления
      </h5>
   </div>
   <div class="formLine">
      <table id="versionsTable" class="customTable table-hover">
         <thead>
            <tr>
               <th class="versionColumn">
                  Номер версии
               </th>
               <th class="dateColumn">
                  Дата выхода
               </th>
               <th class="itsColumn">
                  Диск 1С:ИТС
               </th>
               <th class="previousVersionsColumn">
                  Обновление версии
               </th>
            </tr>
         </thead>
         <tbody>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.18.1363">
                  8.3.18.1363
                  </a>
               </td>
               <td class="dateColumn">
                  16.03.21
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.18.1334">
                  8.3.18.1334
                  </a>
               </td>
               <td class="dateColumn">
                  15.02.21
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.17.1989">
                  8.3.17.1989
                  </a>
               </td>
               <td class="dateColumn">
                  14.02.21
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.18.1289">
                  8.3.18.1289
                  </a>
               </td>
               <td class="dateColumn">
                  12.01.21
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1876">
                  8.3.16.1876
                  </a>
               </td>
               <td class="dateColumn">
                  11.01.21
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.18.1208">
                  8.3.18.1208
                  </a>
               </td>
               <td class="dateColumn">
                  19.11.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.17.1851">
                  8.3.17.1851
                  </a>
               </td>
               <td class="dateColumn">
                  19.11.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1814">
                  8.3.16.1814
                  </a>
               </td>
               <td class="dateColumn">
                  19.11.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.18.1201">
                  8.3.18.1201
                  </a>
               </td>
               <td class="dateColumn">
                  16.11.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1810">
                  8.3.16.1810
                  </a>
               </td>
               <td class="dateColumn">
                  15.11.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr class="success">
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.17.1846">
                  8.3.17.1846
                  </a>
               </td>
               <td class="dateColumn">
                  15.11.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.17.1823">
                  8.3.17.1823
                  </a>
               </td>
               <td class="dateColumn">
                  09.11.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1791">
                  8.3.16.1791
                  </a>
               </td>
               <td class="dateColumn">
                  09.11.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.2107">
                  8.3.15.2107
                  </a>
               </td>
               <td class="dateColumn">
                  09.11.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.18.1128">
                  8.3.18.1128
                  </a>
               </td>
               <td class="dateColumn">
                  12.10.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.17.1549">
                  8.3.17.1549
                  </a>
               </td>
               <td class="dateColumn">
                  07.07.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1659">
                  8.3.16.1659
                  </a>
               </td>
               <td class="dateColumn">
                  06.07.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.17.1496">
                  8.3.17.1496
                  </a>
               </td>
               <td class="dateColumn">
                  04.06.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1985">
                  8.3.15.1985
                  </a>
               </td>
               <td class="dateColumn">
                  03.06.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1502">
                  8.3.16.1502
                  </a>
               </td>
               <td class="dateColumn">
                  03.06.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.17.1386">
                  8.3.17.1386
                  </a>
               </td>
               <td class="dateColumn">
                  23.04.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1958">
                  8.3.15.1958
                  </a>
               </td>
               <td class="dateColumn">
                  22.04.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.14.2095">
                  8.3.14.2095
                  </a>
               </td>
               <td class="dateColumn">
                  22.04.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1359">
                  8.3.16.1359
                  </a>
               </td>
               <td class="dateColumn">
                  22.04.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1296">
                  8.3.16.1296
                  </a>
               </td>
               <td class="dateColumn">
                  31.03.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1224">
                  8.3.16.1224
                  </a>
               </td>
               <td class="dateColumn">
                  26.02.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1148">
                  8.3.16.1148
                  </a>
               </td>
               <td class="dateColumn">
                  20.01.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1869">
                  8.3.15.1869
                  </a>
               </td>
               <td class="dateColumn">
                  19.01.20
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1063">
                  8.3.16.1063
                  </a>
               </td>
               <td class="dateColumn">
                  20.11.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1830">
                  8.3.15.1830
                  </a>
               </td>
               <td class="dateColumn">
                  19.11.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1778">
                  8.3.15.1778
                  </a>
               </td>
               <td class="dateColumn">
                  19.11.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.14.1993">
                  8.3.14.1993
                  </a>
               </td>
               <td class="dateColumn">
                  19.11.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.16.1030">
                  8.3.16.1030
                  </a>
               </td>
               <td class="dateColumn">
                  07.11.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1747">
                  8.3.15.1747
                  </a>
               </td>
               <td class="dateColumn">
                  05.11.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.14.1976">
                  8.3.14.1976
                  </a>
               </td>
               <td class="dateColumn">
                  04.11.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1700">
                  8.3.15.1700
                  </a>
               </td>
               <td class="dateColumn">
                  10.10.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1656">
                  8.3.15.1656
                  </a>
               </td>
               <td class="dateColumn">
                  17.09.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.14.1944">
                  8.3.14.1944
                  </a>
               </td>
               <td class="dateColumn">
                  16.09.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1565">
                  8.3.15.1565
                  </a>
               </td>
               <td class="dateColumn">
                  14.08.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.13.1926">
                  8.3.13.1926
                  </a>
               </td>
               <td class="dateColumn">
                  13.08.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1534">
                  8.3.15.1534
                  </a>
               </td>
               <td class="dateColumn">
                  24.07.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.15.1489">
                  8.3.15.1489
                  </a>
               </td>
               <td class="dateColumn">
                  26.06.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.14.1854">
                  8.3.14.1854
                  </a>
               </td>
               <td class="dateColumn">
                  25.06.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1924">
                  8.3.12.1924
                  </a>
               </td>
               <td class="dateColumn">
                  25.06.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.14.1779">
                  8.3.14.1779
                  </a>
               </td>
               <td class="dateColumn">
                  22.05.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.14.1694">
                  8.3.14.1694
                  </a>
               </td>
               <td class="dateColumn">
                  17.04.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.13.1865">
                  8.3.13.1865
                  </a>
               </td>
               <td class="dateColumn">
                  16.04.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.14.1630">
                  8.3.14.1630
                  </a>
               </td>
               <td class="dateColumn">
                  06.03.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.13.1809">
                  8.3.13.1809
                  </a>
               </td>
               <td class="dateColumn">
                  06.03.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1855">
                  8.3.12.1855
                  </a>
               </td>
               <td class="dateColumn">
                  04.03.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.14.1565">
                  8.3.14.1565
                  </a>
               </td>
               <td class="dateColumn">
                  31.01.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.13.1690">
                  8.3.13.1690
                  </a>
               </td>
               <td class="dateColumn">
                  14.01.19
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.13.1644">
                  8.3.13.1644
                  </a>
               </td>
               <td class="dateColumn">
                  28.11.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1790">
                  8.3.12.1790
                  </a>
               </td>
               <td class="dateColumn">
                  27.11.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.13.1513">
                  8.3.13.1513
                  </a>
               </td>
               <td class="dateColumn">
                  25.09.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1714">
                  8.3.12.1714
                  </a>
               </td>
               <td class="dateColumn">
                  24.09.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1685">
                  8.3.12.1685
                  </a>
               </td>
               <td class="dateColumn">
                  24.09.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1616">
                  8.3.12.1616
                  </a>
               </td>
               <td class="dateColumn">
                  12.09.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1595">
                  8.3.12.1595
                  </a>
               </td>
               <td class="dateColumn">
                  14.08.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1567">
                  8.3.12.1567
                  </a>
               </td>
               <td class="dateColumn">
                  31.07.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1529">
                  8.3.12.1529
                  </a>
               </td>
               <td class="dateColumn">
                  27.06.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1469">
                  8.3.12.1469
                  </a>
               </td>
               <td class="dateColumn">
                  04.06.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1440">
                  8.3.12.1440
                  </a>
               </td>
               <td class="dateColumn">
                  14.05.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.11.3133">
                  8.3.11.3133
                  </a>
               </td>
               <td class="dateColumn">
                  13.05.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2772">
                  8.3.10.2772
                  </a>
               </td>
               <td class="dateColumn">
                  13.05.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.12.1412">
                  8.3.12.1412
                  </a>
               </td>
               <td class="dateColumn">
                  10.04.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.11.3034">
                  8.3.11.3034
                  </a>
               </td>
               <td class="dateColumn">
                  09.02.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2753">
                  8.3.10.2753
                  </a>
               </td>
               <td class="dateColumn">
                  08.02.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.11.2954">
                  8.3.11.2954
                  </a>
               </td>
               <td class="dateColumn">
                  23.01.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.11.2924">
                  8.3.11.2924
                  </a>
               </td>
               <td class="dateColumn">
                  09.01.18
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.11.2899">
                  8.3.11.2899
                  </a>
               </td>
               <td class="dateColumn">
                  12.12.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2699">
                  8.3.10.2699
                  </a>
               </td>
               <td class="dateColumn">
                  12.12.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.11.2867">
                  8.3.11.2867
                  </a>
               </td>
               <td class="dateColumn">
                  21.11.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2667">
                  8.3.10.2667
                  </a>
               </td>
               <td class="dateColumn">
                  14.11.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2650">
                  8.3.10.2650
                  </a>
               </td>
               <td class="dateColumn">
                  31.10.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2639">
                  8.3.10.2639
                  </a>
               </td>
               <td class="dateColumn">
                  20.10.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2580">
                  8.3.10.2580
                  </a>
               </td>
               <td class="dateColumn">
                  22.09.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2561">
                  8.3.10.2561
                  </a>
               </td>
               <td class="dateColumn">
                  11.08.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.9.2309">
                  8.3.9.2309
                  </a>
               </td>
               <td class="dateColumn">
                  11.08.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.2442">
                  8.3.8.2442
                  </a>
               </td>
               <td class="dateColumn">
                  11.08.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2505">
                  8.3.10.2505
                  </a>
               </td>
               <td class="dateColumn">
                  21.07.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2466">
                  8.3.10.2466
                  </a>
               </td>
               <td class="dateColumn">
                  05.07.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2375">
                  8.3.10.2375
                  </a>
               </td>
               <td class="dateColumn">
                  30.06.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2299">
                  8.3.10.2299
                  </a>
               </td>
               <td class="dateColumn">
                  05.06.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.10.2252">
                  8.3.10.2252
                  </a>
               </td>
               <td class="dateColumn">
                  27.04.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.9.2233">
                  8.3.9.2233
                  </a>
               </td>
               <td class="dateColumn">
                  05.04.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.9.2170">
                  8.3.9.2170
                  </a>
               </td>
               <td class="dateColumn">
                  03.02.17
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.9.2033">
                  8.3.9.2033
                  </a>
               </td>
               <td class="dateColumn">
                  19.12.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.2322">
                  8.3.8.2322
                  </a>
               </td>
               <td class="dateColumn">
                  18.12.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.9.1850">
                  8.3.9.1850
                  </a>
               </td>
               <td class="dateColumn">
                  02.11.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.2197">
                  8.3.8.2197
                  </a>
               </td>
               <td class="dateColumn">
                  02.11.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.9.1818">
                  8.3.9.1818
                  </a>
               </td>
               <td class="dateColumn">
                  30.09.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.2167">
                  8.3.8.2167
                  </a>
               </td>
               <td class="dateColumn">
                  28.09.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.2137">
                  8.3.8.2137
                  </a>
               </td>
               <td class="dateColumn">
                  28.09.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.2088">
                  8.3.8.2088
                  </a>
               </td>
               <td class="dateColumn">
                  16.09.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.2054">
                  8.3.8.2054
                  </a>
               </td>
               <td class="dateColumn">
                  31.08.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.2027">
                  8.3.8.2027
                  </a>
               </td>
               <td class="dateColumn">
                  17.08.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.1964">
                  8.3.8.1964
                  </a>
               </td>
               <td class="dateColumn">
                  03.08.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.1933">
                  8.3.8.1933
                  </a>
               </td>
               <td class="dateColumn">
                  22.07.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.1861">
                  8.3.8.1861
                  </a>
               </td>
               <td class="dateColumn">
                  14.07.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.1784">
                  8.3.8.1784
                  </a>
               </td>
               <td class="dateColumn">
                  22.06.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.1747">
                  8.3.8.1747
                  </a>
               </td>
               <td class="dateColumn">
                  22.06.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.1675">
                  8.3.8.1675
                  </a>
               </td>
               <td class="dateColumn">
                  20.05.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.8.1652">
                  8.3.8.1652
                  </a>
               </td>
               <td class="dateColumn">
                  20.04.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.2027">
                  8.3.7.2027
                  </a>
               </td>
               <td class="dateColumn">
                  14.04.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2530">
                  8.3.6.2530
                  </a>
               </td>
               <td class="dateColumn">
                  14.04.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.2008">
                  8.3.7.2008
                  </a>
               </td>
               <td class="dateColumn">
                  25.03.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2524">
                  8.3.6.2524
                  </a>
               </td>
               <td class="dateColumn">
                  25.03.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1993">
                  8.3.7.1993
                  </a>
               </td>
               <td class="dateColumn">
                  23.03.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1970">
                  8.3.7.1970
                  </a>
               </td>
               <td class="dateColumn">
                  18.03.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1949">
                  8.3.7.1949
                  </a>
               </td>
               <td class="dateColumn">
                  04.03.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1917">
                  8.3.7.1917
                  </a>
               </td>
               <td class="dateColumn">
                  10.02.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1901">
                  8.3.7.1901
                  </a>
               </td>
               <td class="dateColumn">
                  04.02.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1873">
                  8.3.7.1873
                  </a>
               </td>
               <td class="dateColumn">
                  28.01.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1860">
                  8.3.7.1860
                  </a>
               </td>
               <td class="dateColumn">
                  19.01.16
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1845">
                  8.3.7.1845
                  </a>
               </td>
               <td class="dateColumn">
                  30.12.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1831">
                  8.3.7.1831
                  </a>
               </td>
               <td class="dateColumn">
                  25.12.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1805">
                  8.3.7.1805
                  </a>
               </td>
               <td class="dateColumn">
                  21.12.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1790">
                  8.3.7.1790
                  </a>
               </td>
               <td class="dateColumn">
                  10.12.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2449">
                  8.3.6.2449
                  </a>
               </td>
               <td class="dateColumn">
                  09.12.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1776">
                  8.3.7.1776
                  </a>
               </td>
               <td class="dateColumn">
                  30.11.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2421">
                  8.3.6.2421
                  </a>
               </td>
               <td class="dateColumn">
                  30.11.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.7.1759">
                  8.3.7.1759
                  </a>
               </td>
               <td class="dateColumn">
                  19.11.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2390">
                  8.3.6.2390
                  </a>
               </td>
               <td class="dateColumn">
                  03.11.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2363">
                  8.3.6.2363
                  </a>
               </td>
               <td class="dateColumn">
                  22.10.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2332">
                  8.3.6.2332
                  </a>
               </td>
               <td class="dateColumn">
                  02.10.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2299">
                  8.3.6.2299
                  </a>
               </td>
               <td class="dateColumn">
                  11.09.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2237">
                  8.3.6.2237
                  </a>
               </td>
               <td class="dateColumn">
                  25.08.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2152">
                  8.3.6.2152
                  </a>
               </td>
               <td class="dateColumn">
                  23.07.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1625">
                  8.3.5.1625
                  </a>
               </td>
               <td class="dateColumn">
                  23.07.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2100">
                  8.3.6.2100
                  </a>
               </td>
               <td class="dateColumn">
                  08.07.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1596">
                  8.3.5.1596
                  </a>
               </td>
               <td class="dateColumn">
                  07.07.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2076">
                  8.3.6.2076
                  </a>
               </td>
               <td class="dateColumn">
                  26.06.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2041">
                  8.3.6.2041
                  </a>
               </td>
               <td class="dateColumn">
                  03.06.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.2014">
                  8.3.6.2014
                  </a>
               </td>
               <td class="dateColumn">
                  21.05.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.1999">
                  8.3.6.1999
                  </a>
               </td>
               <td class="dateColumn">
                  15.05.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1570">
                  8.3.5.1570
                  </a>
               </td>
               <td class="dateColumn">
                  15.05.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.6.1977">
                  8.3.6.1977
                  </a>
               </td>
               <td class="dateColumn">
                  29.04.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1517">
                  8.3.5.1517
                  </a>
               </td>
               <td class="dateColumn">
                  23.03.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1486">
                  8.3.5.1486
                  </a>
               </td>
               <td class="dateColumn">
                  13.03.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1482">
                  8.3.5.1482
                  </a>
               </td>
               <td class="dateColumn">
                  05.03.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1460">
                  8.3.5.1460
                  </a>
               </td>
               <td class="dateColumn">
                  13.02.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1443">
                  8.3.5.1443
                  </a>
               </td>
               <td class="dateColumn">
                  30.01.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1428">
                  8.3.5.1428
                  </a>
               </td>
               <td class="dateColumn">
                  28.01.15
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1383">
                  8.3.5.1383
                  </a>
               </td>
               <td class="dateColumn">
                  12.12.14
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1248">
                  8.3.5.1248
                  </a>
               </td>
               <td class="dateColumn">
                  31.10.14
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1231">
                  8.3.5.1231
                  </a>
               </td>
               <td class="dateColumn">
                  21.10.14
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1186">
                  8.3.5.1186
                  </a>
               </td>
               <td class="dateColumn">
                  03.10.14
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1146">
                  8.3.5.1146
                  </a>
               </td>
               <td class="dateColumn">
                  23.09.14
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1119">
                  8.3.5.1119
                  </a>
               </td>
               <td class="dateColumn">
                  08.08.14
               </td>
               <td class="itsColumn">
                  Сентябрь 2014
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1098">
                  8.3.5.1098
                  </a>
               </td>
               <td class="dateColumn">
                  01.08.14
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1088">
                  8.3.5.1088
                  </a>
               </td>
               <td class="dateColumn">
                  25.07.14
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.5.1068">
                  8.3.5.1068
                  </a>
               </td>
               <td class="dateColumn">
                  10.07.14
               </td>
               <td class="itsColumn">
                  Август 2014
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.4.496">
                  8.3.4.496
                  </a>
               </td>
               <td class="dateColumn">
                  17.06.14
               </td>
               <td class="itsColumn">
                  Июль 2014
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.4.482">
                  8.3.4.482
                  </a>
               </td>
               <td class="dateColumn">
                  29.04.14
               </td>
               <td class="itsColumn">
                  Июнь 2014
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.4.465">
                  8.3.4.465
                  </a>
               </td>
               <td class="dateColumn">
                  04.04.14
               </td>
               <td class="itsColumn">
                  Май 2014
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.4.437">
                  8.3.4.437
                  </a>
               </td>
               <td class="dateColumn">
                  28.02.14
               </td>
               <td class="itsColumn">
                  Апрель 2014
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.4.408">
                  8.3.4.408
                  </a>
               </td>
               <td class="dateColumn">
                  31.01.14
               </td>
               <td class="itsColumn">
                  Март 2014
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.4.389">
                  8.3.4.389
                  </a>
               </td>
               <td class="dateColumn">
                  30.12.13
               </td>
               <td class="itsColumn">
                  Февраль 2014
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.4.365">
                  8.3.4.365
                  </a>
               </td>
               <td class="dateColumn">
                  04.12.13
               </td>
               <td class="itsColumn">
                  Январь 2014
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.3.721">
                  8.3.3.721
                  </a>
               </td>
               <td class="dateColumn">
                  06.09.13
               </td>
               <td class="itsColumn">
                  октябрь 2013
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.3.715">
                  8.3.3.715
                  </a>
               </td>
               <td class="dateColumn">
                  16.08.13
               </td>
               <td class="itsColumn">
                  сентябрь 2013
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.3.687">
                  8.3.3.687
                  </a>
               </td>
               <td class="dateColumn">
                  18.07.13
               </td>
               <td class="itsColumn">
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.3.658">
                  8.3.3.658
                  </a>
               </td>
               <td class="dateColumn">
                  21.06.13
               </td>
               <td class="itsColumn">
                  август 2013
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
            <tr>
               <td class="versionColumn">
                  <a href="/version_files?nick=Platform83&amp;ver=8.3.3.641">
                  8.3.3.641
                  </a>
               </td>
               <td class="dateColumn">
                  29.05.13
               </td>
               <td class="itsColumn">
                  июль 2013
               </td>
               <td class="version previousVersionsColumn">
               </td>
            </tr>
         </tbody>
      </table>
   </div>
   <div class="formLine">
      <a class="project-additional-info" href="http://v8.1c.ru/distrib/">Порядок распространения платформы и прикладных решений (конфигураций) системы программ 1С:Предприятие 8, разрабатываемых фирмой "1С"</a>
   </div>
   <div class="formLine">
      <h5>
         Версии для тестирования
      </h5>
      <p>
         29.01.21 опубликована версия  <a href="/version_files?nick=Platform83&amp;ver=8.3.19.900">8.3.19.900</a>, предназначена для тестирования
      </p>
      <p>
         Предварительные тестовые релизы конфигураций предоставляются партнерам фирмы "1С" и пользователям                 системы программ                 1С:Предприятие для тестирования, предварительного ознакомления с новыми возможностями                 конфигураций, исправлениями ошибок,                 для апробации работы новых релизов на реальных данных.<br>                  Использование предварительного релиза для автоматизации реальных задач предприятия может                 выполняться только в отдельных случаях                 по решению пользователя, совместно с партнером, поддерживающим внедрение
      </p>
   </div>
</div>
`
var releasesTotal = `
<table id="actualTable" class="customTable table-hover">
   <colgroup>
      <col width="340px">
      <col width="80px">
      <col width="140px">
      <col width="80px">
      <col width="140px">
      <col width="140px">
      <col width="86px">
      <col width="140px">
   </colgroup>
   <thead>
      <tr>
         <th class="nameHead" rowspan="2">Название</th>
         <th colspan="2" class="actualVersion">Актуальная версия</th>
         <th colspan="3" class="planVersion">
            Планируемая версия
            <img class="help-icon help-icon-right-th " src="/resources/resources/img/help-icon.png">
            <div class="page-information-dialog page-information-dialog-table-position">
               Указанные сроки являются рабочим планом выпуска (ориентировочным) и могут быть изменены. Не следует указывать приведенные сроки в договорах и других юридически значимых документах. Также не следует использовать данные сроки в каких-либо других формах обязательств (устных, письменных и др.).
            </div>
         </th>
         <th colspan="2" class="testVersion">Версия для ознакомления</th>
      </tr>
      <tr class="second-level-th">
         <th class="versionColumn">Номер версии</th>
         <th class="releaseDate">Дата выхода</th>
         <th class="versionColumn">Номер версии</th>
         <th class="planReleaseDate">Ориентировочная дата выхода</th>
         <th class="updateDate">Дата обновления плановых данных</th>
         <th class="versionColumn">Номер версии</th>
         <th class="publicationDate">Дата публикации</th>
      </tr>
   </thead>
   <tbody>
      <tr group="481">
         <td class="groupColumn" colspan="8">
            <span class="group "></span>
            <span class="group-name">Технологические дистрибутивы</span>
         </td>
      </tr>
      <tr parent-group="481" style="">
         <td class="nameColumn">
            <a href="/project/DevelopmentTools10">1C:Enterprise Development Tools</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=DevelopmentTools10&amp;ver=2020.6.2">2020.6.2</a>
         </td>
         <td class="releaseDate">
            <span>
            10.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn">
            <a href="/version_files?nick=DevelopmentTools10&amp;ver=2021.1 RC2">2021.1 RC2</a>
            <br>
            <a href="/version_files?nick=DevelopmentTools10&amp;ver=2021.1 RC1">2021.1 RC1</a>
            <br>
         </td>
         <td class="publicationDate">
            <span class="">
            09.03.21
            </span><br>
            <span class="">
            26.02.21
            </span><br>
         </td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/Executor">1C:Исполнитель. Бета-версия</a>
         </td>
         <td class="versionColumn"></td>
         <td class="releaseDate"></td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn">
            <a href="/version_files?nick=Executor&amp;ver=2020.2.4.6">2020.2.4.6</a>
            <br>
            <a href="/version_files?nick=Executor&amp;ver=2020.2.3.7">2020.2.3.7</a>
            <br>
            <a href="/version_files?nick=Executor&amp;ver=2020.2.2.31">2020.2.2.31</a>
            <br>
            <a href="/version_files?nick=Executor&amp;ver=2020.2.1.16">2020.2.1.16</a>
            <br>
            <a href="/version_files?nick=Executor&amp;ver=2020.2.0.547">2020.2.0.547</a>
            <br>
         </td>
         <td class="publicationDate">
            <span class="">
            17.02.21
            </span><br>
            <span class="">
            27.11.20
            </span><br>
            <span class="">
            26.10.20
            </span><br>
            <span class="">
            10.08.20
            </span><br>
            <span class="">
            19.06.20
            </span><br>
         </td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/Conversion">1С:Конвертация данных 2.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Conversion&amp;ver=2.1.8.2">2.1.8.2</a>
         </td>
         <td class="releaseDate">
            <span>
            11.06.14
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/Conversion30">1С:Конвертация данных 3</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Conversion30&amp;ver=3.0.5.3">3.0.5.3</a>
         </td>
         <td class="releaseDate">
            <span>
            27.04.17
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn">
            <a href="/version_files?nick=Conversion30&amp;ver=3.1.1.3">3.1.1.3</a>
            <br>
         </td>
         <td class="publicationDate">
            <span class="">
            31.12.20
            </span><br>
         </td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/Translator">1С:Переводчик, редакция 2.1</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Translator&amp;ver=2.1.15.1">2.1.15.1</a>
         </td>
         <td class="releaseDate">
            <span>
            28.04.17
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/BarCode">1С:Печать штрихкодов (ActiveX)</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=BarCode&amp;ver=8.0.16.4">8.0.16.4</a>
         </td>
         <td class="releaseDate">
            <span>
            31.03.14
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/CollaborationSystem">1С:Сервер взаимодействия</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=CollaborationSystem&amp;ver=9.0.33">9.0.33</a>
         </td>
         <td class="releaseDate">
            <span>
            09.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn">
            <a href="/version_files?nick=CollaborationSystem&amp;ver=10.0.37">10.0.37</a>
            <br>
         </td>
         <td class="publicationDate">
            <span class="">
            09.02.21
            </span><br>
         </td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/ScanOpos">1С:Сканер штрихкода (COM)</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ScanOpos&amp;ver=8.1.7.9">8.1.7.9</a>
         </td>
         <td class="releaseDate">
            <span>
            12.04.16
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/STest">1С:Сценарное тестирование 8</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=STest&amp;ver=3.0.24.2">3.0.24.2</a>
         </td>
         <td class="releaseDate">
            <span>
            19.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/Tester">1С:Тестировщик</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Tester&amp;ver=1.0.1.2">1.0.1.2</a>
         </td>
         <td class="releaseDate">
            <span>
            12.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/AddCompDB2">IBM DB2 Express-C</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=AddCompDB2&amp;ver=v9.7 FP6">v9.7 FP6</a>
         </td>
         <td class="releaseDate">
            <span>
            05.10.12
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/AddCompPostgre">PostgreSQL</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=AddCompPostgre&amp;ver=12.5-6.1C">12.5-6.1C</a>
         </td>
         <td class="releaseDate">
            <span>
            17.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn">
            <a href="/version_files?nick=AddCompPostgre&amp;ver=12.6-1.1C">12.6-1.1C</a>
            <br>
            <a href="/version_files?nick=AddCompPostgre&amp;ver=11.11-1.1C">11.11-1.1C</a>
            <br>
            <a href="/version_files?nick=AddCompPostgre&amp;ver=10.16-1.1C">10.16-1.1C</a>
            <br>
         </td>
         <td class="publicationDate">
            <span class=" recentlyChanged">
            15.03.21
            </span><br>
            <span class=" recentlyChanged">
            15.03.21
            </span><br>
            <span class=" recentlyChanged">
            15.03.21
            </span><br>
         </td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/ACC">Автоматизированная проверка конфигураций</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ACC&amp;ver=1.2.6.20">1.2.6.20</a>
         </td>
         <td class="releaseDate">
            <span>
            29.12.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/AddCompDriverHASP">Драйвер аппаратных лицензий платформы 1С:Предприятия (Sentinel HASP)</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=AddCompDriverHASP&amp;ver=7.63">7.63</a>
         </td>
         <td class="releaseDate">
            <span>
            14.05.18
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/ETP">Корпоративный инструментальный пакет 8</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ETP&amp;ver=2.1.7.11">2.1.7.11</a>
         </td>
         <td class="releaseDate">
            <span>
            05.11.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn">
            <a href="/version_files?nick=ETP&amp;ver=2.1.8.2">2.1.8.2</a>
            <br>
         </td>
         <td class="publicationDate">
            <span class="">
            27.01.21
            </span><br>
         </td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/AddCompNetHASP">Менеджер лицензий аппаратной защиты NetHASP</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=AddCompNetHASP&amp;ver=8.31">8.31</a>
         </td>
         <td class="releaseDate">
            <span>
            01.01.07
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/mobile">Мобильная платформа 1С:Предприятия</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=mobile&amp;ver=8.3.18.60">8.3.18.60</a>
         </td>
         <td class="releaseDate">
            <span>
            19.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/TradeWareEpf82">Обработки обслуживания торгового оборудования - для Технологической платформы 8.2</a>
         </td>
         <td class="versionColumn"></td>
         <td class="releaseDate"></td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/Platform80">Технологическая платформа 8.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Platform80&amp;ver=8.0.18.2">8.0.18.2</a>
         </td>
         <td class="releaseDate">
            <span>
            19.12.06
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/Platform81">Технологическая платформа 8.1</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Platform81&amp;ver=8.1.15.14">8.1.15.14</a>
         </td>
         <td class="releaseDate">
            <span>
            30.10.09
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/Platform82">Технологическая платформа 8.2</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Platform82&amp;ver=8.2.19.130">8.2.19.130</a>
         </td>
         <td class="releaseDate">
            <span>
            13.02.15
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/Platform83">Технологическая платформа 8.3</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Platform83&amp;ver=8.3.18.1334">8.3.18.1334</a>
         </td>
         <td class="releaseDate">
            <span>
            15.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn">
            <a href="/version_files?nick=Platform83&amp;ver=8.3.18.1363">8.3.18.1363</a>
            <br>
            <a href="/version_files?nick=Platform83&amp;ver=8.3.19.900">8.3.19.900</a>
            <br>
         </td>
         <td class="publicationDate">
            <span class="">
            06.03.21
            </span><br>
            <span class="">
            29.01.21
            </span><br>
         </td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/PlatformUtil">Утилита администрирования конфигураций и информационных баз 1С:Предприятия 8</a>
         </td>
         <td class="versionColumn"></td>
         <td class="releaseDate"></td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn">
            <a href="/version_files?nick=PlatformUtil&amp;ver=8.2.13.219">8.2.13.219</a>
            <br>
         </td>
         <td class="publicationDate">
            <span class="">
            06.04.11
            </span><br>
         </td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/EnterpriseLicenseTools">Утилита лицензирования 1С:Предприятия (1C:Enterprise License Tools)</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=EnterpriseLicenseTools&amp;ver=0.15.0.2">0.15.0.2</a>
         </td>
         <td class="releaseDate">
            <span>
            16.11.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="481" style="
         ">
         <td class="nameColumn">
            <a href="/project/ReportFactory">Фабрика отчетов</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ReportFactory&amp;ver=1.0.1.1">1.0.1.1</a>
         </td>
         <td class="releaseDate">
            <span>
            26.04.12
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr group="473">
         <td class="groupColumn" colspan="8">
            <span class="group "></span>
            <span class="group-name">Стандартные библиотеки</span>
         </td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/DMIL">1С:Библиотека интеграции с 1С:Документооборотом</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=DMIL&amp;ver=1.1.17.2">1.1.17.2</a>
         </td>
         <td class="releaseDate">
            <span>
            26.08.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/MDLP">1С:Библиотека интеграции с МДЛП</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=MDLP&amp;ver=1.2.3.5">1.2.3.5</a>
         </td>
         <td class="releaseDate">
            <span>
            17.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/ISL21">1С:Библиотека интернет-поддержки пользователей, редакция 2.1</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ISL21&amp;ver=2.1.9.18">2.1.9.18</a>
         </td>
         <td class="releaseDate">
            <span>
            08.02.18
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/ISL22">1С:Библиотека интернет-поддержки пользователей, редакция 2.2</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ISL22&amp;ver=2.2.3.19">2.2.3.19</a>
         </td>
         <td class="releaseDate">
            <span>
            29.01.19
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/ISL23">1С:Библиотека интернет-поддержки пользователей, редакция 2.3</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ISL23&amp;ver=2.3.4.18">2.3.4.18</a>
         </td>
         <td class="releaseDate">
            <span>
            19.06.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/ISL24">1С:Библиотека интернет-поддержки пользователей, редакция 2.4</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ISL24&amp;ver=2.4.2.63">2.4.2.63</a>
         </td>
         <td class="releaseDate">
            <span>
            15.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/ISL25">1С:Библиотека интернет-поддержки пользователей, редакция 2.5</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ISL25&amp;ver=2.5.1.31">2.5.1.31</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            12.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/MobileCEL">1С:Библиотека подключаемого оборудования для мобильных приложений</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=MobileCEL&amp;ver=2.13.7.0">2.13.7.0</a>
         </td>
         <td class="releaseDate">
            <span>
            11.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/CEL21">1С:Библиотека подключаемого оборудования, редакция 2.1</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=CEL21&amp;ver=2.1.4.14">2.1.4.14</a>
         </td>
         <td class="releaseDate">
            <span>
            20.01.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/CEL30">1С:Библиотека подключаемого оборудования, редакция 3.0</a>
         </td>
         <td class="versionColumn"></td>
         <td class="releaseDate"></td>
         <td class="versionColumn ">
            <a href="/project/CEL30#plan-release-link-3.0.2"> 3.0.2
            <br>
            </a><a href="/project/CEL30#plan-release-link-3.0.1"> 3.0.1
            <br>
            </a>
         </td>
         <td class="planReleaseDate">
            <span class="">
            02.09.19
            </span><br>
            <span class="">
            11.03.19
            </span><br>
         </td>
         <td class="updateDate">
            <span class="">
            28.06.19
            </span><br>
            <span class="">
            20.11.18
            </span><br>
         </td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SSL">1С:Библиотека стандартных подсистем 8.2</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SSL&amp;ver=2.0.1.19">2.0.1.19</a>
         </td>
         <td class="releaseDate">
            <span>
            26.07.12
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SSL21">1С:Библиотека стандартных подсистем 8.2, редакция 2.1</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SSL21&amp;ver=2.1.9.2">2.1.9.2</a>
         </td>
         <td class="releaseDate">
            <span>
            06.08.14
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SSL22">1С:Библиотека стандартных подсистем, редакция 2.2</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SSL22&amp;ver=2.2.5.36">2.2.5.36</a>
         </td>
         <td class="releaseDate">
            <span>
            14.07.15
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SSL23">1С:Библиотека стандартных подсистем, редакция 2.3</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SSL23&amp;ver=2.3.7.10">2.3.7.10</a>
         </td>
         <td class="releaseDate">
            <span>
            02.04.18
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SSL24">1С:Библиотека стандартных подсистем, редакция 2.4</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SSL24&amp;ver=2.4.6.241">2.4.6.241</a>
         </td>
         <td class="releaseDate">
            <span>
            21.12.18
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SSL30">1С:Библиотека стандартных подсистем, редакция 3.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SSL30&amp;ver=3.0.3.341">3.0.3.341</a>
         </td>
         <td class="releaseDate">
            <span>
            04.07.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SSL31">1С:Библиотека стандартных подсистем, редакция 3.1</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SSL31&amp;ver=3.1.4.174">3.1.4.174</a>
         </td>
         <td class="releaseDate">
            <span>
            05.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SMTL11">1С:Библиотека технологии сервиса, редакция 1.1</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SMTL11&amp;ver=1.1.4.5">1.1.4.5</a>
         </td>
         <td class="releaseDate">
            <span>
            21.10.19
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SMTL12">1С:Библиотека технологии сервиса, редакция 1.2</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SMTL12&amp;ver=1.2.2.35">1.2.2.35</a>
         </td>
         <td class="releaseDate">
            <span>
            30.07.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/SMTL20">1С:Библиотека технологии сервиса, редакция 2.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SMTL20&amp;ver=2.0.4.39">2.0.4.39</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            10.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/LED">1С:Библиотека электронных документов, редакция 1.1</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=LED&amp;ver=1.1.28.35">1.1.28.35</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            12.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/LED12">1С:Библиотека электронных документов, редакция 1.2</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=LED12&amp;ver=1.2.7.11">1.2.7.11</a>
         </td>
         <td class="releaseDate">
            <span>
            07.08.15
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/LED13">1С:Библиотека электронных документов, редакция 1.3</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=LED13&amp;ver=1.3.11.125">1.3.11.125</a>
         </td>
         <td class="releaseDate">
            <span>
            15.02.19
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/LED14">1С:Библиотека электронных документов, редакция 1.4</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=LED14&amp;ver=1.4.1.75">1.4.1.75</a>
         </td>
         <td class="releaseDate">
            <span>
            05.02.19
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/LED15">1С:Библиотека электронных документов, редакция 1.5</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=LED15&amp;ver=1.5.1.93">1.5.1.93</a>
         </td>
         <td class="releaseDate">
            <span>
            09.08.19
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/LED16">1С:Библиотека электронных документов, редакция 1.6</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=LED16&amp;ver=1.6.4.117">1.6.4.117</a>
         </td>
         <td class="releaseDate">
            <span>
            24.07.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/LED17">1С:Библиотека электронных документов, редакция 1.7</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=LED17&amp;ver=1.7.2.109">1.7.2.109</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            11.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/LED18">1С:Библиотека электронных документов, редакция 1.8</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=LED18&amp;ver=1.8.1.50">1.8.1.50</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            11.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="473" style="
         ">
         <td class="nameColumn">
            <a href="/project/LED19">1С:Библиотека электронных документов, редакция 1.9</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=LED19&amp;ver=1.9.1.26">1.9.1.26</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            11.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr group="4">
         <td class="groupColumn" colspan="8">
            <span class="group "></span>
            <span class="group-name">Типовые конфигурации фирмы "1С" для России </span>
         </td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/Treasury275">1С:275ФЗ</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Treasury275&amp;ver=1.0.2.12">1.0.2.12</a>
         </td>
         <td class="releaseDate">
            <span>
            19.05.16
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/StorekeeperDevelopers">1С:Кладовщик. Версия для разработчиков</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=StorekeeperDevelopers&amp;ver=1.0.18.1">1.0.18.1</a>
         </td>
         <td class="releaseDate">
            <span>
            29.04.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/StorekeeperIts">1С:Кладовщик для ИТС</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=StorekeeperIts&amp;ver=1.0.18">1.0.18</a>
         </td>
         <td class="releaseDate">
            <span>
            29.04.20
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/ClientEDO82">1С:Клиент ЭДО 8, редакция 1.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=ClientEDO82&amp;ver=1.0.3.8">1.0.3.8</a>
         </td>
         <td class="releaseDate">
            <span>
            29.05.15
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/CashdeskDev">1С:Мобильная касса (для разработчиков)</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=CashdeskDev&amp;ver=3.8.15.0">3.8.15.0</a>
         </td>
         <td class="releaseDate">
            <span>
            27.01.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn">
            <a href="/version_files?nick=CashdeskDev&amp;ver=3.1.5.0">3.1.5.0</a>
            <br>
            <a href="/version_files?nick=CashdeskDev&amp;ver=3.0.22.0">3.0.22.0</a>
            <br>
         </td>
         <td class="publicationDate">
            <span class="">
            25.12.19
            </span><br>
            <span class="">
            26.09.19
            </span><br>
         </td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/CorporatePerformanceManagement13">1С:Управление холдингом 1.3</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=CorporatePerformanceManagement13&amp;ver=1.3.12.1">1.3.12.1</a>
         </td>
         <td class="releaseDate">
            <span>
            08.10.18
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/CorporatePerformanceManagement30">1С:Управление холдингом 3.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=CorporatePerformanceManagement30&amp;ver=3.0.18.2">3.0.18.2</a>
         </td>
         <td class="releaseDate">
            <span>
            09.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/CorporatePerformanceManagement31">1С:Управление холдингом 3.1</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=CorporatePerformanceManagement31&amp;ver=3.1.7.2">3.1.7.2</a>
         </td>
         <td class="releaseDate">
            <span>
            11.02.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/AccountingCorp">Бухгалтерия предприятия КОРП, редакция 2.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=AccountingCorp&amp;ver=2.0.66.135">2.0.66.135</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            12.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/AccountingCorp30">Бухгалтерия предприятия КОРП, редакция 3.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=AccountingCorp30&amp;ver=3.0.89.51">3.0.89.51</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            12.03.21
            </span>
         </td>
         <td class="versionColumn ">
            <a href="/project/AccountingCorp30#plan-release-link-3.0.91"> 3.0.91
            <br>
            </a><a href="/project/AccountingCorp30#plan-release-link-3.0.90"> 3.0.90
            <br>
            </a>
         </td>
         <td class="planReleaseDate">
            <span class="">
            Апрель 2021
            </span><br>
            <span class="">
            Март 2021
            </span><br>
         </td>
         <td class="updateDate">
            <span class=" recentlyChanged">
            22.03.21
            </span><br>
            <span class=" recentlyChanged">
            22.03.21
            </span><br>
         </td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/Accounting20_82">Бухгалтерия предприятия, редакция 2.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Accounting20_82&amp;ver=2.0.66.135">2.0.66.135</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            12.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/Accounting30">Бухгалтерия предприятия, редакция 3.0</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=Accounting30&amp;ver=3.0.89.51">3.0.89.51</a>
         </td>
         <td class="releaseDate">
            <span class="recentlyChanged">
            12.03.21
            </span>
         </td>
         <td class="versionColumn ">
            <a href="/project/Accounting30#plan-release-link-3.0.91"> 3.0.91
            <br>
            </a><a href="/project/Accounting30#plan-release-link-3.0.90"> 3.0.90
            <br>
            </a>
         </td>
         <td class="planReleaseDate">
            <span class="">
            Апрель 2021
            </span><br>
            <span class="">
            Март 2021
            </span><br>
         </td>
         <td class="updateDate">
            <span class=" recentlyChanged">
            22.03.21
            </span><br>
            <span class="">
            22.02.21
            </span><br>
         </td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/HRMCorp">Зарплата и управление персоналом КОРП, редакция 2.5</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=HRMCorp&amp;ver=2.5.159.4">2.5.159.4</a>
         </td>
         <td class="releaseDate">
            <span>
            05.03.21
            </span>
         </td>
         <td class="versionColumn ">
            <a href="/project/HRMCorp#plan-release-link-2.5."> 2.5.
            <br>
            </a>
         </td>
         <td class="planReleaseDate">
            <span class="">
            1 кв. 2021
            </span><br>
         </td>
         <td class="updateDate">
            <span class="">
            05.03.21
            </span><br>
         </td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr parent-group="4" style="
         ">
         <td class="nameColumn">
            <a href="/project/HRMCorp30">Зарплата и управление персоналом КОРП, редакция 3</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=HRMCorp30&amp;ver=3.1.14.436">3.1.14.436</a>
         </td>
         <td class="releaseDate">
            <span>
            04.03.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
      <tr group="162">
         <td class="groupColumn" colspan="8">
            <span class="group "></span>
            <span class="group-name">Отраслевые решения</span>
         </td>
      </tr>
      <tr parent-group="162" style="
         ">
         <td class="nameColumn">
            <a href="/project/SLK">1С:СЛК</a>
         </td>
         <td class="versionColumn actualVersionColumn">
            <a href="/version_files?nick=SLK&amp;ver=3.0.24.9152">3.0.24.9152</a>
         </td>
         <td class="releaseDate">
            <span>
            27.01.21
            </span>
         </td>
         <td class="versionColumn"></td>
         <td class="planReleaseDate"><span class="noDate">Не определена</span></td>
         <td class="updateDate"></td>
         <td class="versionColumn"></td>
         <td class="publicationDate"><span class="noDate">Не определена</span></td>
      </tr>
   </tbody>
</table>
`
