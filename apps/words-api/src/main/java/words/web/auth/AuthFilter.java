package words.web.auth;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

import javax.inject.Inject;
import javax.servlet.*;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;
import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.util.Base64;
import java.util.Optional;

import static org.springframework.http.HttpStatus.UNAUTHORIZED;

@Slf4j
@Component
public class AuthFilter implements Filter {

    private static final String USERNAME = "username";

    private final AuthService authService;

    @Inject
    public AuthFilter(AuthService authService) {
        this.authService = authService;
    }

    @Override
    public void init(FilterConfig filterConfig) throws ServletException {
    }

    @Override
    public void destroy() {
    }

    @Override
    public void doFilter(ServletRequest servletRequest, ServletResponse servletResponse, FilterChain filterChain) throws IOException, ServletException {
        HttpServletRequest request = (HttpServletRequest) servletRequest;
        HttpServletResponse response = (HttpServletResponse) servletResponse;

        HttpSession session = request.getSession(true);

        String userFromSession = (String) session.getAttribute(USERNAME);
        if (userFromSession != null) {
            filterChain.doFilter(servletRequest, servletResponse);
            return;
        }

        Optional<String> user = getAuthorizedUsername(request.getHeader("Authorization"));
        if (!user.isPresent()) {
            response.sendError(UNAUTHORIZED.value(), UNAUTHORIZED.getReasonPhrase());
            return;
        }

        session.setAttribute(USERNAME, user.get());
        filterChain.doFilter(servletRequest, servletResponse);
    }

    private Optional<String> getAuthorizedUsername(String header) {
        if (header == null) {
            log.error("auth header is null");
            return Optional.empty();
        }
        try {
            String[] tokens = extractAndDecodeHeader(header);
            String username = tokens[0];
            String password = tokens[1];

            if (authService.isUserExist(username, password)) {
                return Optional.of(username);
            }

        } catch (Exception e) {
            log.error("auth failed", e);
        }
        return Optional.empty();
    }

    private String[] extractAndDecodeHeader(String header) throws UnsupportedEncodingException {
        byte[] base64Token = header.substring(6).getBytes("UTF-8");
        byte[] decoded = Base64.getDecoder().decode(base64Token);

        String token = new String(decoded, "UTF-8");
        int delim = token.indexOf(":");

        return new String[] {token.substring(0, delim), token.substring(delim + 1)};
    }

}
